package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
	"github.com/suhlig/plaintweet/server"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves a plain-text representation of a single tweet via HTTP",
	RunE: func(command *cobra.Command, args []string) error {
		port, found := os.LookupEnv("PORT")

		if !found {
			port = "8080"
		}

		server := server.NewServer(plaintweet.NewRepository(command.Context())).WithBlurb(command.Short)

		maxUpTimeStr, found := os.LookupEnv("MAX_UPTIME")

		if found {
			m, err := time.ParseDuration(maxUpTimeStr)

			if err != nil {
				return err
			}

			server = server.WithMaxUptime(m)
		}

		return server.Start(":" + port)
	},
}
