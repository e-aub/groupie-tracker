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
	var wg global.CheckWG
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
	wg.WG.Add(4)
	go global.Read(w, r, artist_url, &context.Artists,&wg)
	go global.Read(w, r, locations_url, &context.Locations,&wg)
	go global.Read(w, r, dates_url, &context.Dates,&wg)
	go global.Read(w, r, relations_url, &context.Relations,&wg)
	pages := []string{
		"template/pages/details.html",
	}
	wg.WG.Wait()
	global.ExecuteTemplate(w, r, pages, context)
}
