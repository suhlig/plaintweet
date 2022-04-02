package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleLiveness(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)
	w.Header().Add("Server", plaintweet.VersionStringShort())
	w.Header().Add("X-Uptime", time.Since(s.startTime).Round(time.Second).String())

	if s.maxUptime != nil && time.Since(s.startTime) > *s.maxUptime {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: Maximum uptime of %v reached\n", s.maxUptime)
	} else {
		fmt.Fprintln(w, "OK")
	}
}
