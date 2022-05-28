package plaintweet

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type TwitterRepository struct {
	twitter *twitter.Client
}

func NewRepository(ctx context.Context, consumerKey, consumerSecret string) Repository {
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	return &TwitterRepository{twitter: twitter.NewClient(config.Client(ctx))}
}

func (r *TwitterRepository) Lookup(id int64) (PlainTweet, error) {
	tweet, _, err := r.twitter.Statuses.Show(id, &twitter.StatusShowParams{TweetMode: "extended"})

	return &V1Tweet{tweet}, err
}

func (r *TwitterRepository) Find(uri *url.URL) (PlainTweet, error) {
	id, err := findTweetID(uri)

	if err != nil {
		return nil, err
	}

	return r.Lookup(id)
}

type V1Tweet struct {
	Tweet *twitter.Tweet
}

func (pt V1Tweet) String() string {
	text := plaintext(pt.Tweet.FullText).
		replaceHtmlEntities().
		unQuote()

	var on string

	if pt.Tweet.Lang == "de" {
		on = "auf"
	} else {
		on = "on"
	}

	return fmt.Sprintf(`"%s" -- @%s %s #twitter `, text, pt.Tweet.User.ScreenName, on)
}
