package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		fmt.Printf("Path=%q RawPath=%q key=%q\n", r.URL.Path, r.URL.RawPath, key)
		fmt.Fprintf(w, "Path=%q RawPath=%q key=%q\n", r.URL.Path, r.URL.RawPath, key)
	})
	http.ListenAndServe(":3333", r)
}
