package plaintweet

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/dghubble/go-twitter/twitter"
)

type MockRepository struct {
	id   int64
	text string
	user string
}

func NewMockRepository(id int64, t, u string) Repository {
	return &MockRepository{id: id, text: t, user: u}
}

func (*MockRepository) AuthenticatedUser() (string, error) {
	return "mock", nil
}

func (r *MockRepository) Lookup(id int64) (*PlainTweet, error) {
	if id == r.id {
		return &PlainTweet{Tweet: &twitter.Tweet{FullText: r.text, Lang: "en", User: &twitter.User{ScreenName: r.user}}}, nil
	} else {
		return nil, errors.New("no status found with that ID")
	}
}

func (r *MockRepository) Find(uri *url.URL) (*PlainTweet, error) {
	if uri.Path == fmt.Sprintf("/%d", r.id) {
		return &PlainTweet{Tweet: &twitter.Tweet{FullText: r.text, Lang: "en", User: &twitter.User{ScreenName: r.user}}}, nil
	} else {
		return nil, errors.New("no status found with that ID")
	}
}
