package plaintweet

import (
	"fmt"
	"html"
	"net/url"
	"path"
	"strconv"
)

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

func findTweetID(uri *url.URL) (int64, error) {
	_, idStr := path.Split(uri.Path)

	if idStr == "" {
		return 0, fmt.Errorf("cannot find a tweet ID in %s", uri)
	}

	return strconv.ParseInt(idStr, 10, 64)
}
