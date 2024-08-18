package handlers

import (
	"groupie_tracker/global"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	var artists []global.Artist
	if r.URL.Path != "/" {
		http.Error(w, "page not found!", 404)
		return
	}
	url := "/artists"
	global.Read(w, r, url, &artists)
	pages := []string{
		"template/pages/home.html",
	}
	global.ExecuteTemplate(w, r, pages, artists)
}
