package global

import (
	"encoding/json"
	"net/http"
)

var BaseUrl string = "https://groupietrackers.herokuapp.com/api"

func Read(w http.ResponseWriter, err chan error, url string, data any, wg *CheckWG) {
	if !wg.NotWG {
		defer wg.WG.Done()
	}
	res, res_err := http.Get(BaseUrl + url)
	if res_err != nil {
		err <- res_err
		return

	}
	defer res.Body.Close()
	json_err := json.NewDecoder(res.Body).Decode(data)
	if json_err != nil {
		err <- json_err
		return
	}
}
