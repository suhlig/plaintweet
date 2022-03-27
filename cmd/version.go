package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(plaintweet.VersionString())
	},
}
