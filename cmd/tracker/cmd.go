package main

import (
	"github.com/spf13/cobra"
)

const (
	DefaultConfDir = "/etc/tracker"

	DefaultAppDir = "/var/lib/tracker"
)

var rootCmd = &cobra.Command{
	Use:   "tracker [command]",
	Short: "your command description",
	Long:  `your command description`,
}

var (
	Author    string
	Version   string
	BuildTime string
)

func init() {
	// subcommands
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(serverCmd)
}

func main() {
	rootCmd.Execute()
}
