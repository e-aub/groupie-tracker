package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/global"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		global.HandleError(w, r, global.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed!"})
		return
	}
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
	var err error
	_, err = strconv.Atoi(id)
	if err != nil || r.URL.Path != ("/artists/"+id) {
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "page not found!"})
		return
	}

	artist_url := "/artists/" + id
	locations_url := "/locations/" + id
	dates_url := "/dates/" + id
	relations_url := "/relation/" + id
	// get data from api
	wg.WG.Add(4)
	go global.Read(w, &err, artist_url, &context.Artists, &wg)
	go global.Read(w, &err, locations_url, &context.Locations, &wg)
	go global.Read(w, &err, dates_url, &context.Dates, &wg)
	go global.Read(w, &err, relations_url, &context.Relations, &wg)
	if err != nil {
		global.HandleError(w, r, global.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	if context.Artists.Id == 0 {
		
		global.HandleError(w, r, global.Error{Code: http.StatusNoContent, Message: "No Content!"})
		return
	}

	wg.WG.Wait()
	pages := []string{
		"template/pages/details.html",
	}

	global.ExecuteTemplate(w, r, pages, context)
}
