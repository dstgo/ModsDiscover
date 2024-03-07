package main

import (
	"context"
	"github.com/dstgo/task"
	"github.com/dstgo/tracker/internal/conf"
	"github.com/dstgo/tracker/internal/core/server"
	"github.com/dstgo/tracker/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log/slog"
	"os/signal"
	"path"
	"syscall"
)

var (
	configFile string
	initial    bool
)

var serverCmd = &cobra.Command{
	Use:          "server [-f filename]",
	Short:        "Run backend server",
	Example:      "myapp server -f /etc/myapp/config.yaml",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := serve(configFile, Author, Version, BuildTime)
		if err != nil && !errors.Is(err, context.Canceled) {
			slog.Error(errors.Wrap(err, "server running failed").Error())
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(&configFile, "conf", "f", path.Join(DefaultConfDir, "config.yaml"), "specified server config file")
	serverCmd.Flags().BoolVar(&initial, "i", false, "only initial server data, not run web server")
}

func newServer(ctx context.Context, configFile string, author string, version string, buildTime string) (*server.HttpServer, error) {

	// read configuration
	appConfig := config.NewConfigFile(configFile)
	if err := appConfig.ReadConfig(); err != nil {
		return nil, err
	}

	// map configuration struct
	appConf, err := conf.NewAppConf(appConfig, conf.BuildInfo{
		Author:    author,
		Version:   version,
		BuildTime: buildTime,
	})
	if err != nil {
		return nil, err
	}

	// ini logger
	logger, err := server.NewLogger(appConf.LogConf)
	if err != nil {
		return nil, err
	}

	// set app mode
	gin.SetMode(gin.ReleaseMode)

	// initialize app server
	app, err := server.NewHTTPApp(
		server.WithCtx(ctx),
		server.WithConf(appConf),
		server.WithLogger(logger),
	)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func serve(configFile string, author string, version string, buildTime string) error {

	// listen signal
	signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	// create app
	app, err := newServer(signalCtx, configFile, author, version, buildTime)
	if err != nil {
		return err
	}

	serverTask, causeFunc := task.New(signalCtx)
	defer causeFunc(nil)
	defer app.Shutdown()

	serverWorker := task.NewWorker(func(ctx context.Context) error {
		if initial {
			return nil
		}
		// run the http server
		return app.Run()
	})

	serverTask.Add(serverWorker)

	return serverTask.Run()
}
