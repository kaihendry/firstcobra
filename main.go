/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/kaihendry/firstcobra/cmd"
	"golang.org/x/exp/slog"
)

func main() {
	// JSON logging is the default
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
	cmd.Execute()
}
