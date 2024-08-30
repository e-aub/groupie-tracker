package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupie_tracker/global"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancelCause(r.Context())

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

	artist_url := "https://groupietrackers.herokuapp.com/api/artists/" + id
	locations_url := "https://groupietrackers.herokuapp.com/api/locations/" + id
	dates_url := "https://groupietrackers.herokuapp.com/api/dates/" + id
	relations_url := "https://groupietrackers.herokuapp.com/api/relation/" + id
	// get data from api
	errchan := make(chan error)
	done := make(chan struct{})

	wg.Add(4)
	go func() {
		global.FetchGoRoutine(ctx, errchan, locations_url, &context.Locations, &wg, "locations")
		global.GetLocationsId(ctx, &context.Locations, errchan)
		if err != nil {
			errchan <- err
			return
		}
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
		cancel(err)
		fmt.Println(err)
		global.HandleError(w, r, global.Error{Code: http.StatusNotFound, Message: "not found!"})
	case <-done:
		// If done without errors, proceed to execute the template
		pages := []string{"template/pages/details.html"}
		global.ExecuteTemplate(w, r, pages, context)
	}
}
