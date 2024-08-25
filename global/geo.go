package global

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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

type (
	Index struct {
		Location1 []Location `json:"index"`
	}

	Location struct {
		Location []string `json:"locations"`
	}

	GeoResponse struct {
		Results []struct {
			PlaceID string `json:"place_id"`
		} `json:"results"`
		Status string `json:"status"`
	}
)

func GetLocationsId(data *ArtistLocation) error {
	apiKey := "AIzaSyBOypms8DmfMpEx6-IRJzwz7lvBmE4kr94"
	for _, location := range data.Locations {
		var res GeoResponse
		adress := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
		err := Fetch(adress, &res)
		if err != nil {
			// errChan <- err
			return err
		} else if res.Status != "OK" {
			// errChan <- err
			return err
		}
		*&data.LocationsIds = append(*&data.LocationsIds, res.Results[0].PlaceID)

	}

	// defer func() {
	// 	wg.Done()
	// }()
	return nil
}
