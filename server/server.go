package server

import (
	"log"
	"net/http"
	"time"

	"github.com/suhlig/plaintweet/plaintweet"
)

type Server struct {
	startTime time.Time
	maxUptime *time.Duration
	blurb     string
}

func NewServer() *Server {
	return &Server{} // Syntactic sugar for the fluent interface
}

func (s *Server) WithMaxUptime(maxUpTime time.Duration) *Server {
	s.maxUptime = &maxUpTime
	return s
}

func (s *Server) WithBlurb(blurb string) *Server {
	s.blurb = blurb
	return s
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/liveness", s.HandleLiveness)   // The kubelet uses liveness probes to know when to restart a container
	http.HandleFunc("/readiness", s.HandleReadiness) // The kubelet uses readiness probes to know when a container is ready to start accepting traffic
	// TODO http.HandleFunc("/startup", s.HandleStartup) // The kubelet uses startup probes to know when a container application has started.
	http.HandleFunc("/version", s.HandleVersion)
	http.HandleFunc("/", s.HandleRoot)

	log.Printf("Starting server %s on port %s", plaintweet.VersionStringShort(), addr)

	if s.maxUptime != nil {
		log.Printf("Maximum allowed uptime set to %v; afterwards /liveness will report an error", s.maxUptime)
	}

	s.startTime = time.Now()
	return http.ListenAndServe(addr, nil)
}
