package global

import (
	"encoding/json"
	"net/http"
	"sync"
)

var BaseUrl string = "https://groupietrackers.herokuapp.com/api"

func Read(err chan error, url string, data any, wg *sync.WaitGroup) {
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
	defer wg.Done()
}
