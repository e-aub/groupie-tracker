package handlers

import (
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
	url := "/artists"
	var wg global.CheckWG
	var err error
	wg.NotWG = true
	global.Read(w, &err, url, &artists, &wg)
	if err != nil {
		global.HandleError(w, r, global.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	pages := []string{
		"template/pages/home.html",
		"template/components/carousel.html",
	}
	global.ExecuteTemplate(w, r, pages, artists)
}
