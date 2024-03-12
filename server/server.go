package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dstgo/lobbyapi"
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/api"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/data"
	hertzslog "github.com/hertz-contrib/logger/slog"
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

	// load logger
	closer, err := loadLogger(appConf.Log)
	if err != nil {
		return nil, err
	}

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
	hertz, err := newHttpServer(appConf.Http, redisDB)
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

// returns a new logger
func loadLogger(logConf conf.LogConf) (io.Closer, error) {
	dir := filepath.Dir(logConf.File)
	if dir != "." && len(dir) > 0 {
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return nil, err
		}
	}

	// open log file
	logFile, err := os.OpenFile(logConf.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	// set level
	var level = new(slog.LevelVar)
	if err := level.UnmarshalText([]byte(logConf.Level)); err != nil {
		return nil, err
	}

	// set logger
	hlog.SetLogger(hertzslog.NewLogger(
		hertzslog.WithLevel(level),
		hertzslog.WithOutput(writer),
	))

	return logFile, nil
}
