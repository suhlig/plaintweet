package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
)

var RootCmd = &cobra.Command{
	Use:   "plaintweet",
	Short: "Provides a plain-text representation of a single tweet",
}

func selectRepo(ctx context.Context) (plaintweet.Repository, error) {
	token, isSet := os.LookupEnv("TWITTER_BEARER_TOKEN")

	if isSet {
		// prefer V2 if possible
		return plaintweet.NewV2Repository(ctx, token), nil
	} else {
		// fall back to V1
		consumerKey, isSet := os.LookupEnv("TWITTER_CONSUMER_KEY")

		if !isSet {
			return nil, fmt.Errorf("neither TWITTER_BEARER_TOKEN nor TWITTER_CONSUMER_KEY are set")
		}

		consumerSecret, isSet := os.LookupEnv("TWITTER_CONSUMER_SECRET")

		if !isSet {
			return nil, fmt.Errorf("TWITTER_CONSUMER_SECRET is missing, but required when using TWITTER_CONSUMER_KEY")
		}

		return plaintweet.NewRepository(ctx, consumerKey, consumerSecret), nil
	}
}
