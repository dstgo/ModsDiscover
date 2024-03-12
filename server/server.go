package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/dstgo/lobbyapi"
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/api"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/data"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func NewTracker(ctx context.Context, file string) (*server.Hertz, error) {
	// parse config
	appConf, err := conf.Load(file)
	if err != nil {
		return nil, err
	}

	logger, closer, err := newLogger(appConf.Log)
	if err != nil {
		return nil, err
	}
	slog.SetDefault(logger)

	// load mongodb
	mgodb, err := data.LoadMongoDB(ctx, appConf.DB)
	if err != nil {
		return nil, err
	}
	// load redis
	redisDB, err := data.LoadRedisDB(ctx, appConf.Redis)
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

	if err := api.Register(hertz, mgodb, redisDB, lobbyClient, steamClient); err != nil {
		return nil, err
	}

	// on shutdown
	hertz.OnShutdown = append(hertz.OnShutdown, func(ctx context.Context) {
		if err := mgodb.Close(ctx); err != nil {
			slog.Error("failed to close mongodb client", err)
		}
		if err := redisDB.Close(); err != nil {
			slog.Error("failed to close redis client", err)
		}
		_ = closer.Close()
	})

	return hertz, nil
}

// returns a new hertz http server
func newHttpServer(httpConf conf.HttpConf) (*server.Hertz, error) {
	hertz := server.New(
		server.WithHostPorts(httpConf.Listen),
		server.WithReadTimeout(httpConf.ReadTimeout),
		server.WithIdleTimeout(httpConf.IdleTimeout),
		server.WithBasePath(httpConf.BasePath),
	)

	return hertz, nil
}

// returns a new logger
func newLogger(logConf conf.LogConf) (*slog.Logger, io.Closer, error) {
	dir := filepath.Dir(logConf.File)
	if dir != "." && len(dir) > 0 {
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return nil, nil, err
		}
	}

	// open log file
	logFile, err := os.OpenFile(logConf.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	// set level
	var level slog.Level
	if err := level.UnmarshalText([]byte(logConf.Level)); err != nil {
		return nil, nil, err
	}

	// new logger
	logger := slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	}))
	return logger, nil, nil
}
