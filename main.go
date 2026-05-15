package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var (
	Version = "dev"
	Commit  = "N/A"
)

const message = "Hello world"

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = fmt.Fprint(w, message)
	})
	log.Printf("version: %s, commit: %s", Version, Commit)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
