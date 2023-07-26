package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "firstcobra",
	RunE: func(cmd *cobra.Command, args []string) error {
		log, err := cmd.PersistentFlags().GetString("log")
		if err != nil {
			return err
		}
		if log == "text" {
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
		}
		slog.Info("root called", "log", log)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("failed to execute", "err", err)
		os.Exit(1)
	}
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	// check for -log switch, if -log text is passed, use text logging
	rootCmd.PersistentFlags().StringP("log", "l", "json", "Log format: json or text")
}
