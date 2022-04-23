package plaintweet

import (
	"net/url"
)

type Repository interface {
	// returns the authenticated user's screen name or an error if the authentication was not ok
	AuthenticatedUser() (string, error)

	// returns the PlainTweet identified by its id
	Lookup(id int64) (*PlainTweet, error)

	// returns the PlainTweet identified by its URL
	Find(uri *url.URL) (*PlainTweet, error)
}
