package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"groupie_tracker/global"
)

func TestGeo(t *testing.T) {
	var locations global.Locations
	err := global.Fetch("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		t.Error(err)
	}
	for _, locationsGroup := range locations.Index {
		for _, location := range locationsGroup.Locations {
			var response global.GeoResponse
			url := fmt.Sprintf("https://maps-data.p.rapidapi.com/geocoding.php?query=%s&lang=en&country=fr", location)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Error(err)
			}

			req.Header.Add("x-rapidapi-key", "3deb8955ecmsh5bfc47a8a9587b3p19cb4ajsn0a1824e4a63c")
			req.Header.Add("x-rapidapi-host", "maps-data.p.rapidapi.com")

			res, err := http.DefaultClient.Do(req)
			if res.StatusCode != 200 || err != nil {
				t.Error("not 200 or error")
			}
			defer res.Body.Close()

			err = json.NewDecoder(res.Body).Decode(&response)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(response)

		}
	}
}
