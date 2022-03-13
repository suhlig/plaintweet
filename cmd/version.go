package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ldflags will be set by goreleaser
var version = "vDEV"
var commit = "NONE"
var date = "UNKNOWN"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("plaintweet %s (%s), built on %s\n", version, commit, date)
	},
}
