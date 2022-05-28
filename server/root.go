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
		log.Printf("%v : 200", r.URL)
		fmt.Fprintln(w, s.blurb)
		return
	}

	tweet, err := s.repository.Find(r.URL)

	if err != nil {
		w.WriteHeader(404)
		log.Printf("%v : 404 : %s", r.URL, err)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	log.Printf("%v : 200 : %s", r.URL, tweet)
	fmt.Fprintf(w, "%s\n", tweet)
}
