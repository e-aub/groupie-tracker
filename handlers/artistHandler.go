package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"groupie_tracker/global"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if r.Method != http.MethodGet {
		global.HandleError(w, r, global.MethodNotAllowedErr)
		return
	}
	var (
		baseUrl        = "https://groupietrackers.herokuapp.com/api/"
		artistDetailes struct {
			Artist    global.Artist
			Locations global.ArtistLocation
			Dates     global.ArtistDate
			Relations global.ArtistRelation
		}
		wg sync.WaitGroup
	)

	// handle url
	id := strings.TrimPrefix(r.URL.Path, "/artists/")
	// var err error
	if !global.IsId(id) {
		global.HandleError(w, r, global.NotFoundErr)
		return
	}
	artistUrl := baseUrl + "artists/" + id
	locationsUrl := baseUrl + "locations/" + id
	datesUrl := baseUrl + "dates/" + id
	relationsUrl := baseUrl + "relation/" + id
	// get data from api
	errchan := make(chan error)
	done := make(chan struct{})

	wg.Add(4)
	go func() {
		global.FetchGoRoutine(ctx, errchan, locationsUrl, &artistDetailes.Locations, &wg, "locations")
		global.GetLocationsId(ctx, &artistDetailes.Locations, errchan)
		defer wg.Done()
	}()
	go global.FetchGoRoutine(ctx, errchan, artistUrl, &artistDetailes.Artist, &wg, "")
	go global.FetchGoRoutine(ctx, errchan, datesUrl, &artistDetailes.Dates, &wg, "")
	go global.FetchGoRoutine(ctx, errchan, relationsUrl, &artistDetailes.Relations, &wg, "")

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
		global.HandleError(w, r, global.NotFoundErr)
	case <-done:
		// If done without errors, proceed to execute the template
		pages := []string{"template/pages/details.html"}
		global.ExecuteTemplate(w, r, pages, artistDetailes)
	}
}
