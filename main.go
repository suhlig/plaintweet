package main

import (
	"context"
	"fmt"
	"os"

	"github.com/suhlig/plaintweet/cmd"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(cmd.ServeCmd)
	rootCmd.AddCommand(cmd.PrintCmd)
	rootCmd.AddCommand(cmd.VersionCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
