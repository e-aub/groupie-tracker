package global

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

var BaseUrl string = "https://groupietrackers.herokuapp.com/api"

func Read(ctx context.Context, err chan error, url string, data any, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !strings.HasPrefix(url, "/locations") {
				defer wg.Done()
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
			return
		}
	}

}
