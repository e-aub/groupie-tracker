package handlers

import (
	"groupie_tracker/global"
	"net/http"
	"sync"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	var artists []global.Artist
	// if r.URL.Path != "/" {
	// 	http.Error(w, "page not found!", 404)
	// 	return
	// }
	url := "/artists"
	var wg sync.WaitGroup
	wg.Add(1)
	go global.Read(w, r, url, &artists, &wg)
	pages := []string{
		"template/pages/home.html",
	}
	wg.Wait()
	global.ExecuteTemplate(w, r, pages, artists)
}
