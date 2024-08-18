package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/global"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	var context struct {
		Artists   global.Artist
		Locations global.ArtistLocation
		Dates     global.ArtistDate
		Relations global.ArtistRelation
	}
	// handle url
	url_path := strings.Split(r.URL.Path, "/")
	id := url_path[2]

	_, err := strconv.Atoi(id)
	if err != nil || r.URL.Path != ("/artists/"+id) {
		http.Error(w, "page not found", 404)
		return
	}
	artist_url := "/artists/" + id
	locations_url := "/locations/" + id
	dates_url := "/dates/" + id
	relations_url := "/relation/" + id
	// get data from api
	global.Read(w, r, artist_url, &context.Artists)
	global.Read(w, r, locations_url, &context.Locations)
	global.Read(w, r, dates_url, &context.Dates)
	global.Read(w, r, relations_url, &context.Relations)
	pages := []string{
		"template/pages/details.html",
	}
	global.ExecuteTemplate(w, r, pages, context)
}
