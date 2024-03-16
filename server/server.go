package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/api"
	"github.com/dstgo/tracker/assets"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/internal/data"
	"github.com/dstgo/tracker/internal/jobs"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/pkg/lobbyapi"
	"os/signal"
	"syscall"
	"time"

	// embedded timezone data
	_ "time/tzdata"
)

type Server struct {
	server  *server.Hertz
	timeout time.Duration
	cleanup func(ctx context.Context)
}

func (s *Server) Serve() {
	ctx := context.Background()

	notifyContext, cancelFunc := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	defer cancelFunc()

	run := make(chan error)

	// run the server
	go func() {
		s.server.Spin()
		run <- s.server.Run()
		close(run)
	}()

	select {
	case <-notifyContext.Done():
		hlog.Infof("receive close signal, ready to shutdown, max timmeout %s", s.timeout)
	case err := <-run:
		hlog.Infof("running failed: error=%v", err)
	}

	// close the server
	if err := s.server.Close(); err != nil {
		hlog.Infof("server closed failed: error=%v", err)
	}

	// shutdown
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	done := make(chan struct{})

	go func() {
		s.cleanup(timeoutCtx)
		done <- struct{}{}
		close(done)
	}()

	select {
	case <-timeoutCtx.Done():
		hlog.Error("shutdown timeout")
	case <-done:
		hlog.Info("shutdown finished")
	}
}

func NewTracker(ctx context.Context, logger hlog.FullLogger, appConf *conf.AppConf) (*Server, error) {
	// load mongodb
	mgodb, err := data.LoadMongoDB(ctx, appConf.DB)
	if err != nil {
		return nil, err
	}

	// new server
	hertz, err := newHttpServer(appConf.Http)
	if err != nil {
		return nil, err
	}

	// dst api initial
	lobbyClient := lobbyapi.New(appConf.Dst.KleiToken)
	steamClient, err := steamapi.New(appConf.Dst.SteamKey)
	if err != nil {
		return nil, err
	}

	// geoip db
	geoIpDB, err := data.LoadGeoIpDBInMem(assets.GeopIp2CityDB)
	if err != nil {
		return nil, err
	}

	// app environment
	env := &types.Env{
		Logger:   logger,
		MongoDB:  mgodb,
		LobbyCLI: lobbyClient,
		SteamCLI: steamClient,
		GeoIpDB:  geoIpDB,
	}

	// register api router
	apis, err := api.NewRouter(ctx, hertz, env)
	if err != nil {
		return nil, err
	}

	// load cron jobs
	cronJobs, err := jobs.LoadCronJobs(appConf.Dst, apis.Lobby.LobbyHandler)
	if err != nil {
		return nil, err
	}

	// after started
	hertz.OnRun = append(hertz.OnRun, func(ctx context.Context) error {
		// run all jobs
		go cronJobs.Run()
		return nil
	})

	// on shutdown
	onShutdown := func(ctx context.Context) {
		// wait for all jobs were stopped
		<-cronJobs.Stop().Done()
		hlog.Info("cron jobs stopped successfully")

		if err := mgodb.Close(context.Background()); err != nil {
			hlog.Error("failed to close mongodb client", err)
		}
		hlog.Info("mongodb closed successfully")

		hlog.Info("redis closed successfully")

		if err := geoIpDB.Close(); err != nil {
			hlog.Error("failed to close geo db", err)
		}
		hlog.Info("geodb closed successfully")
	}

	return &Server{
		server:  hertz,
		timeout: time.Second * 5,
		cleanup: onShutdown,
	}, nil
}
