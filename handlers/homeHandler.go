package handlers

import (
	"fmt"
	"net/http"

	"groupie_tracker/global"
)

var (
	artists []global.Artist
	pages   = []string{
		"template/pages/home.html",
		"template/components/carousel.html",
	}
	url = "https://groupietrackers.herokuapp.com/api/artists"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		global.HandleError(w, r, global.MethodNotAllowedErr)
		return
	}
	if r.URL.Path != "/" {
		global.HandleError(w, r, global.NotFoundErr)
		return
	}

	err := global.Fetch(url, &artists)
	if err != nil {
		fmt.Println(err)
		global.HandleError(w, r, global.InternalServerErr)
		return
	}

	global.ExecuteTemplate(w, r, pages, artists)
}
