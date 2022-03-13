package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "plaintweet",
	Short: "Provides a plain-text representation of a single tweet",
}

func main() {
	rootCmd.AddCommand(cmd.ServeCmd)
	rootCmd.AddCommand(cmd.PrintCmd)
	rootCmd.AddCommand(cmd.VersionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
