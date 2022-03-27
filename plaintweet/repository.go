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

type Repository struct {
	twitter *twitter.Client
}

func NewRepository(ctx context.Context) *Repository {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	return &Repository{twitter: twitter.NewClient(
		config.Client(
			ctx,
		),
	),
	}
}

// returns the authenticated user's screen name or an error if the authentication was not ok
func (r *Repository) AuthenticatedUser() (string, error) {
	response, _, err := r.twitter.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})

	if err != nil {
		return "", err
	}

	return response.ScreenName, nil
}

// returns the PlainTweet identified by its id
func (r *Repository) Lookup(id int64) (*PlainTweet, error) {
	tweet, _, err := r.twitter.Statuses.Show(id, &twitter.StatusShowParams{TweetMode: "extended"})

	return &PlainTweet{tweet}, err
}

// returns the PlainTweet identified by its URL
func (r *Repository) Find(uri *url.URL) (*PlainTweet, error) {
	_, idStr := path.Split(uri.Path)
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
