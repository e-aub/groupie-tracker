package global

import (
	"errors"
	"fmt"
)

func GetLocationsId(data *ArtistLocation, errChan chan error) {
	apiKey := "AIzaSyBOypms8DmfMpEx6-IRJzwz7lvBmE4kr94"
	for _, location := range data.Locations {
		var res GeoResponse
		adress := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
		err := Fetch(adress, &res)
		if err != nil {
			errChan <- err
			return
		} else if res.Status != "OK" {
			errChan <- errors.New(res.Status)
			return
		}
		data.LocationsIds = append(data.LocationsIds, fmt.Sprintf("%g,%g", res.Results[0].Geometry.Location.Lat, res.Results[0].Geometry.Location.Lng))
	}
}
