package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "firstcobra",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("failed to execute", "err", err)
		os.Exit(1)
	}
	// check for -log switch, if -log text is passed, use text logging
	log, err := rootCmd.Flags().GetString("log")
	if err != nil {
		slog.Error("failed to get log flag", "err", err)
	}
	if log == "text" {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	}
	slog.Info("parsed args", "log", log)
}

func init() {
	// check for -log switch, if -log text is passed, use text logging
	rootCmd.Flags().StringP("log", "l", "json", "Log format: json or text")
}
