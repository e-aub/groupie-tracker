package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"groupie_tracker/global"
)

var baseUrl = "https://groupietrackers.herokuapp.com/api/"

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

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
	id := strings.TrimPrefix(r.URL.Path, "/artists/")
	// var err error
	if !global.IsId(id) {
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "page not found!"})
		return
	}
	artist_url := baseUrl + "artists/" + id
	locations_url := baseUrl + "locations/" + id
	dates_url := baseUrl + "dates/" + id
	relations_url := baseUrl + "relation/" + id
	// get data from api
	errchan := make(chan error)
	done := make(chan struct{})

	wg.Add(4)
	go func() {
		global.FetchGoRoutine(ctx, errchan, locations_url, &context.Locations, &wg, "locations")
		global.GetLocationsId(ctx, &context.Locations, errchan)
		defer wg.Done()
	}()
	go global.FetchGoRoutine(ctx, errchan, artist_url, &context.Artists, &wg, "")
	go global.FetchGoRoutine(ctx, errchan, dates_url, &context.Dates, &wg, "")
	go global.FetchGoRoutine(ctx, errchan, relations_url, &context.Relations, &wg, "")

	go func() {
		wg.Wait()
		close(done)
		close(errchan)
	}()

	// Listen for the first error or completion

	select {
	case err := <-errchan:
		fmt.Println(err)
		cancel()
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "not found!"})
	case <-done:
		// If done without errors, proceed to execute the template
		pages := []string{"template/pages/details.html"}
		global.ExecuteTemplate(w, r, pages, context)
	}
}
