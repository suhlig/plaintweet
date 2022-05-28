package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", plaintweet.VersionStringShort())

	_, err := s.repository.Lookup(20)
	var status string

	if err == nil {
		status = "OK"
	} else {
		status = fmt.Sprintf("Error: %v", err)
		w.WriteHeader(500)
	}

	log.Printf("%v : %s", r.URL, status)
	fmt.Fprintln(w, status)
}
