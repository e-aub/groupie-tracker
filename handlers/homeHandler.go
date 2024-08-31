package handlers

import (
	"fmt"
	"net/http"

	"groupie_tracker/global"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		global.HandleError(w, r, global.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed!"})
		return
	}
	if r.URL.Path != "/" {
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "page not found!"})
		return
	}

	var artists []global.Artist
	url := "https://groupietrackers.herokuapp.com/api/artists"

	err := global.Fetch(url, &artists)
	if err != nil {
		fmt.Println(err)
		global.HandleError(w, r, global.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	pages := []string{
		"template/pages/home.html",
		"template/components/carousel.html",
	}
	global.ExecuteTemplate(w, r, pages, artists)
}
