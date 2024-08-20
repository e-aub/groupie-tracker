package main

import (
	"groupie_tracker/global"
	"groupie_tracker/handlers"
	"net/http"
)

func main() {
	port := ":8000"
	mux := handlers.CreateRouter()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	global.CreateServer(mux, port)
}
