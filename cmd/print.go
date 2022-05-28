package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var PrintCmd = &cobra.Command{
	Use:   "print TWEET_ID | URL",
	Short: "Prints a plain-text representation of a single tweet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a tweet ID or URL")
		}

		_, err := url.Parse(args[0])

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true // no need to print usage; we'll handle all errors

		uri, err := url.Parse(args[0])

		if err != nil {
			return err
		}

		repo, err := selectRepo(cmd.Context())

		if err != nil {
			return err
		}

		tweet, err := repo.Find(uri)

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", tweet)

		return nil
	},
}
