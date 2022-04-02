package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhlig/plaintweet/plaintweet"
)

func (s *Server) HandleVersion(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)
	w.Header().Add("Server", plaintweet.VersionStringShort())
	fmt.Fprintf(w, "%s\n", plaintweet.VersionString())
}
