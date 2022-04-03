package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleLiveness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", plaintweet.VersionStringShort())
	w.Header().Add("X-Uptime", time.Since(s.startTime).Round(time.Second).String())

	var status string

	if s.maxUptime == nil || time.Since(s.startTime) < *s.maxUptime {
		status = "OK"
	} else {
		w.WriteHeader(500)
		status = fmt.Sprintf("Error: Maximum uptime of %v reached\n", s.maxUptime)
	}

	log.Printf("%v: %s", r.URL, status)
	fmt.Fprintln(w, status)
}
