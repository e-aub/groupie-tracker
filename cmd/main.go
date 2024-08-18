package main

import (
	"groupie_tracker/global"
	"groupie_tracker/handlers"
)

func main() {
	port := ":8000"
	mux := handlers.CreateRouter()
	global.CreateServer(mux, port)
}
