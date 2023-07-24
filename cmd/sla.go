/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

// slaCmd represents the sla command
var slaCmd = &cobra.Command{
	Use:   "sla",
	Short: "Accept the service level agreement",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("sla called")
	},
}

func init() {
	rootCmd.AddCommand(slaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
