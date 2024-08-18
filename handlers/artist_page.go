package handlers

import (
	"groupie_tracker/global"
	"net/http"
	"strconv"
	"strings"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	var artist global.Artist
	url_path := strings.Split(r.URL.Path, "/")
	id := url_path[2]
	_, err := strconv.Atoi(id)
	if err != nil || r.URL.Path != ("/artists/"+id) {
		http.Error(w, "page not found", 404)
		return
	}
	url := "/artists/" + id
	global.Read(w, r, url, &artist)
	pages := []string{
		"template/pages/details.html",
	}
	global.ExecuteTemplate(w, r, pages, artist)
}
