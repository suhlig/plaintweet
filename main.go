package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/suhlig/plaintweet/plaintweet"
	"golang.org/x/oauth2/clientcredentials"
)

// ldflags will be set by goreleaser
var version = "vDEV"
var commit = "NONE"
var date = "UNKNOWN"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Missing argument for tweet (URL or ID)")
		os.Exit(MissingArgument)
	}

	if os.Args[1] == "version" {
		fmt.Printf("%s %s (%s), built on %s\n", getProgramName(), version, commit, date)
		return
	}

	uri, err := url.Parse(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(ParseError)
	}

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	repo := plaintweet.NewRepository(
		twitter.NewClient(
			config.Client(
				context.Background(),
			),
		),
	)

	tweet, err := repo.Find(uri)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(NoSuchTweet)
	}

	fmt.Printf("%s\n", tweet)
}

func getProgramName() string {
	path, err := os.Executable()

	if err != nil {
		os.Stderr.WriteString("Warning: Could not determine program name; using 'unknown'.\n")
		return "unknown"
	}

	return filepath.Base(path)
}
