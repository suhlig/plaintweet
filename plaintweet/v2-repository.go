package plaintweet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type V2Repository struct {
	client *twitter.Client
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func NewV2Repository(ctx context.Context, token string) Repository {
	return &V2Repository{
		client: &twitter.Client{
			Authorizer: authorize{
				Token: token,
			},
			Client: http.DefaultClient,
			Host:   "https://api.twitter.com",
		},
	}
}

func (r *V2Repository) Lookup(id int64) (PlainTweet, error) {
	opts := twitter.TweetLookupOpts{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{twitter.TweetFieldLanguage, twitter.TweetFieldText},
	}

	idstr := fmt.Sprintf("%d", id)
	tweetResponse, err := r.client.TweetLookup(context.Background(), []string{idstr}, opts)

	if err != nil {
		return nil, err
	}

	if tweetResponse.Raw.Errors != nil {
		return nil, errors.New(tweetResponse.Raw.Errors[0].Detail)
	}

	if dictionary, ok := tweetResponse.Raw.TweetDictionaries()[idstr]; !ok {
		return nil, fmt.Errorf("requested tweet was not returned by Twitter")
	} else {
		return &V2Tweet{
			Text:     dictionary.Tweet.Text,
			Language: dictionary.Tweet.Language,
			UserName: dictionary.Author.UserName,
		}, nil
	}
}

func (r *V2Repository) Find(uri *url.URL) (PlainTweet, error) {
	id, err := findTweetID(uri)

	if err != nil {
		return nil, err
	}

	return r.Lookup(id)
}

type V2Tweet struct {
	UserName string
	Text     string
	Language string
}

func (t V2Tweet) String() string {
	text := plaintext(t.Text).
		replaceHtmlEntities().
		unQuote()

	var on string

	if t.Language == "de" {
		on = "auf"
	} else {
		on = "on"
	}

	return fmt.Sprintf(`"%s" -- @%s %s #twitter `, text, t.UserName, on)
}
