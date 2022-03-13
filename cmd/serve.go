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
	RunE: func(cmd *cobra.Command, args []string) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.URL)

			tweet, err := plaintweet.NewRepository(r.Context()).Find(r.URL)

			if err != nil {
				w.WriteHeader(404)
				fmt.Fprintf(w, "%s\n", err)
				return
			}

			fmt.Fprintf(w, "%s\n", tweet)
		})

		port, found := os.LookupEnv("PORT")

		if !found {
			port = "8080"
		}

		log.Printf("Starting server on port %s", port)

		return http.ListenAndServe(":"+port, nil)
	},
}
