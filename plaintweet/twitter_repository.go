package plaintweet

import (
	"context"
	"fmt"
	"html"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type TwitterRepository struct {
	twitter *twitter.Client
}

func NewRepository(ctx context.Context) Repository {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	return &TwitterRepository{twitter: twitter.NewClient(
		config.Client(
			ctx,
		),
	),
	}
}

func (r *TwitterRepository) AuthenticatedUser() (string, error) {
	response, _, err := r.twitter.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})

	if err != nil {
		return "", err
	}

	return response.ScreenName, nil
}

func (r *TwitterRepository) Lookup(id int64) (*PlainTweet, error) {
	tweet, _, err := r.twitter.Statuses.Show(id, &twitter.StatusShowParams{TweetMode: "extended"})

	return &PlainTweet{tweet}, err
}

func (r *TwitterRepository) Find(uri *url.URL) (*PlainTweet, error) {
	_, idStr := path.Split(uri.Path)

	if idStr == "" {
		return nil, fmt.Errorf("cannot find a tweet ID in %s", uri)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return nil, err
	}

	return r.Lookup(id)
}

type PlainTweet struct {
	Tweet *twitter.Tweet
}

func (pt PlainTweet) String() string {
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

type plaintext string

func (pt plaintext) replaceHtmlEntities() plaintext {
	return plaintext(html.UnescapeString(string(pt)))
}

func (pt plaintext) unQuote() plaintext {
	unquoted, err := strconv.Unquote(string(pt))

	if err != nil {
		return pt
	}

	return plaintext(unquoted)
}
