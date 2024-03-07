package main

import (
	"github.com/dstgo/filebox"
	"github.com/dstgo/tracker/assets"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log/slog"
)

var genDir string
var genCmd = &cobra.Command{
	Use:          "gen [-d dir]",
	Short:        "Generate the default config directory",
	Long:         "Generate the default config directory, if already exists, all files will be overwritten.",
	SilenceUsage: true,
	Example:      "wilson gen -d /etc/wilson",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateResourceDir(genDir)
		if err != nil {
			slog.Error(err.Error())
		}
	},
}

func init() {
	genCmd.Flags().StringVarP(&genDir, "dest", "d", DefaultConfDir, "default resource directory")
}

func generateResourceDir(dir string) error {
	err := filebox.CopyFs(assets.Fs, "wilson", dir)
	if err != nil {
		return errors.Wrap(err, "generate config failed")
	}
	return nil
}
