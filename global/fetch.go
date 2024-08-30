package global

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
)

func FetchGoRoutine(ctx context.Context, err chan error, url string, data any, wg *sync.WaitGroup, purpose string) {
	select {
	case <-ctx.Done():
		return
	default:
		if purpose != "locations" {
			defer wg.Done()
		}
		res, res_err := http.Get(url)
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
}

func Fetch(url string, typ any) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	} else if response.StatusCode != 200 {
		return errors.New("not 200")
	}
	defer response.Body.Close()
	Data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(Data, typ)
	if err != nil {
		return err
	}
	return nil
}
