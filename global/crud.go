package global

import (
	"encoding/json"
	"net/http"
)

var BaseUrl string = "https://groupietrackers.herokuapp.com/api"

func Read(w http.ResponseWriter, r *http.Request, url string, data any, wg *CheckWG) {
	if !wg.NotWG {
		defer wg.WG.Done()
	}
	res, err := http.Get(BaseUrl + url)
	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(data)
}
