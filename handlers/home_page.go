package handlers

import (
	"net/http"
	"sync"

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
	// url := "/artists"
	var wg sync.WaitGroup
	// wg.NotWG = true
	errchan := make(chan error)
	done := make(chan bool)
	wg.Add(1)

	// go global.Read(errchan, url, &artists, &wg)

	go func() {
		wg.Wait()
		close(done)
		close(errchan)
	}()

	select {
	case err := <-errchan:
		// Handle the first error and return
		global.HandleError(w, r, global.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	case <-done:
		pages := []string{
			"template/pages/home.html",
			"template/components/carousel.html",
		}
		global.ExecuteTemplate(w, r, pages, artists)
	}
}
