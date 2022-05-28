package plaintweet

import (
	"net/url"
)

type PlainTweet interface {
	String() string
}

type Repository interface {
	// returns the PlainTweet identified by its id
	Lookup(id int64) (PlainTweet, error)

	// returns the PlainTweet identified by its URL
	Find(uri *url.URL) (PlainTweet, error)
}
