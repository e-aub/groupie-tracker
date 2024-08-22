package global

import (
	"fmt"
	"log"
	"net/http"
)

func CreateServer(mux *http.ServeMux, port string) {
	fmt.Printf("\x1b[0;32mServer Listening in ->\x1b[1;32m http://localhost%s\x1b[0;39m\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		log.Fatal(server_err.Error())
	}
}
