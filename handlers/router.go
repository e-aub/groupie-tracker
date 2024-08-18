package handlers

import (
	"net/http"
)

func CreateRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/artists/{id}", ArtistPage)
	return mux
}
