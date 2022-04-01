package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhlig/plaintweet/plaintweet"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves a plain-text representation of a single tweet via HTTP",
	RunE: func(command *cobra.Command, args []string) error {
		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())
			fmt.Fprintln(w, "OK")
		})

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

		http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())
			fmt.Fprintf(w, "%s\n", plaintweet.VersionString())
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)
			w.Header().Add("Server", plaintweet.VersionStringShort())

			if r.URL.Path == "/" {
				fmt.Fprintln(w, command.Root().Short)
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

		return http.ListenAndServe(":"+port, nil)
	},
}
