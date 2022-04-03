package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", plaintweet.VersionStringShort())

	if r.URL.Path == "/" {
		log.Printf("%v: OK", r.URL)
		fmt.Fprintln(w, s.blurb)
		return
	}

	tweet, err := plaintweet.NewRepository(r.Context()).Find(r.URL)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	log.Printf("%v: %s", r.URL, tweet)
	fmt.Fprintf(w, "%s\n", tweet)
}
