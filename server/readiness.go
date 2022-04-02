package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)
	w.Header().Add("Server", plaintweet.VersionStringShort())

	_, err := plaintweet.NewRepository(r.Context()).Lookup(20)

	if err == nil {
		fmt.Fprintln(w, "OK")
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: %s\n", err)
	}
}
