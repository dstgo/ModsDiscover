package main

import (
	"context"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/server"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:          "tracker [command]",
	Long:         "dst information collector server,\nfor more information to access https://github.com/dstgo/tracker.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// parse configuration
		appConf, err := conf.Load(configFile)
		if err != nil {
			return err
		}

		// load logger
		logger, closer, err := server.NewLogger(appConf.Log)
		if err != nil {
			return err
		}
		defer closer.Close()

		tracker, err := server.NewTracker(ctx, logger, appConf)
		if err != nil {
			return err
		}
		tracker.Serve()
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "/etc/tracker/config.yaml", "tracker config file")
}

func main() {
	rootCmd.Execute()
}
