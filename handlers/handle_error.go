package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Error struct {
	Message string
	Code    int
}

// HandleError sends an HTTP response with the specified error code and renders an error page using the provided error details.
func HandleError(w http.ResponseWriter, r *http.Request, errType Error) {
	w.WriteHeader(errType.Code)
	// http.Error(w, errType.Message, errType.Code)
	tmpl, err := template.ParseFiles("template/pages/error.html")
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "error", errType)
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
}
