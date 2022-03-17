package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chiyutianyi/gittools/pkg/version"
)

// GitindexCmd represents the base command when called without any subcommands
var GitindexCmd = &cobra.Command{
	Use:     "gittools",
	Short:   "git tools",
	Version: version.Version(),
}

func main() {
	if err := GitindexCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "start gittools error: %v", err)
		os.Exit(1)
	}
}
