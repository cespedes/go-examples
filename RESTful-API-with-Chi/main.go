package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := chi.NewRouter()
	r.Get("/albums", getAlbums)
	r.Get("/albums/{id}", getAlbumByID)
	r.Post("/albums", postAlbums)
	http.ListenAndServe(":3333", r)
}

// outIndentedJSON writes to w an indented form of the JSON-encoded v.
func outIndentedJSON(w http.ResponseWriter, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(w)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(w http.ResponseWriter, r *http.Request) {
	outIndentedJSON(w, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum album
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newAlbum); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	outIndentedJSON(w, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			outIndentedJSON(w, a)
			return
		}
	}
	http.Error(w, `{"message": "album not found"}`, http.StatusNotFound)
}
