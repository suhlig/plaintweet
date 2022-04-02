package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)
	w.Header().Add("Server", plaintweet.VersionStringShort())

	if r.URL.Path == "/" {
		fmt.Fprintln(w, s.blurb)
		return
	}

	tweet, err := plaintweet.NewRepository(r.Context()).Find(r.URL)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "%s\n", tweet)
}
