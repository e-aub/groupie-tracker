package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

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
	var wg sync.WaitGroup

	// handle url
	url_path := strings.Split(r.URL.Path, "/")
	id := url_path[2]
	// var err error
	_, err := strconv.Atoi(id)
	if err != nil || r.URL.Path != ("/artists/"+id) {
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "page not found!"})
		return
	}

	artist_url := "/artists/" + id
	locations_url := "/locations/" + id
	dates_url := "/dates/" + id
	relations_url := "/relation/" + id
	// get data from api
	errchan := make(chan error)
	done := make(chan bool)

	wg.Add(4)
	go global.Read(errchan, artist_url, &context.Artists, &wg)
	go global.Read(errchan, locations_url, &context.Locations, &wg)
	go global.Read(errchan, dates_url, &context.Dates, &wg)
	go global.Read(errchan, relations_url, &context.Relations, &wg)

	go func() {
		wg.Wait()
		close(done)
		close(errchan)
	}()

	// Listen for the first error or completion
	select {
	case <-errchan:
		// Handle the first error and return
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "not found!"})
		return
	case <-done:
		global.GetLocationsId(&context.Locations)
		// If done without errors, proceed to execute the template
		pages := []string{"template/pages/details.html"}
		global.ExecuteTemplate(w, r, pages, context)
	}
}
