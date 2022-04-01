package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
)

var startTime time.Time
var maxUpTime *time.Duration

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves a plain-text representation of a single tweet via HTTP",
	RunE: func(command *cobra.Command, args []string) error {
		// The kubelet uses liveness probes to know when to restart a container
		http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())
			w.Header().Add("X-Uptime", time.Since(startTime).String())

			if maxUpTime != nil && time.Since(startTime) > *maxUpTime {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Error: Maximum uptime of %v reached\n", maxUpTime)
			} else {
				fmt.Fprintln(w, "OK")
			}
		})

		// The kubelet uses readiness probes to know when a container is ready to start accepting traffic
		http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())

			_, err := plaintweet.NewRepository(r.Context()).Lookup(20)

			if err == nil {
				fmt.Fprintln(w, "OK")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Error: %s\n", err)
			}
		})

		// TODO The kubelet uses startup probes to know when a container application has started.

		http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())
			fmt.Fprintf(w, "%s\n", plaintweet.VersionString())
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())

			if r.URL.Path == "/" {
				fmt.Fprintln(w, command.Short)
				return
			}

			tweet, err := plaintweet.NewRepository(r.Context()).Find(r.URL)

			if err != nil {
				w.WriteHeader(404)
				fmt.Fprintf(w, "Error: %s\n", err)
				return
			}

			fmt.Fprintf(w, "%s\n", tweet)
		})

		port, found := os.LookupEnv("PORT")

		if !found {
			port = "8080"
		}

		log.Printf("Starting server %s on port %s", plaintweet.VersionStringShort(), port)
		startTime = time.Now()
		setMaxUptime()

		return http.ListenAndServe(":"+port, nil)
	},
}

func setMaxUptime() {
	maxUpTimeStr, found := os.LookupEnv("MAX_UPTIME")

	if found {
		m, err := time.ParseDuration(maxUpTimeStr)

		if err != nil {
			panic(err)
		}

		maxUpTime = &m
		log.Printf("Maximum allowed uptime set to %v; afterwards /liveness will report an error", maxUpTime)
	}
}
