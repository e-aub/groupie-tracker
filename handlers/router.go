package handlers

import (
	"net/http"
)

func CreateRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/artists/{id}", ArtistPage)
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))
	return mux
}
