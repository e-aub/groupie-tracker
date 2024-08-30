package main

import (
	"fmt"
	"groupie_tracker/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/artists/{id}", handlers.ArtistPage)
	fmt.Printf("\x1b[0;32mServer Listening in ->\x1b[1;32m http://localhost%s\x1b[0;39m\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
