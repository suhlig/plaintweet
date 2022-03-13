package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
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
		uri, err := url.Parse(args[0])

		if err != nil {
			return err
		}

		tweet, err := plaintweet.NewRepository(context.Background()).Find(uri)

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", tweet)

		return nil
	},
}
