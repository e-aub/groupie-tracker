package global

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func GetLocationsId(ctx context.Context, data *ArtistLocation, errChan chan error) {
	var (
		geoWg sync.WaitGroup
		mutex sync.Mutex
	)
	data.LocationsCoordinates = map[string]string{}
	select {
	case <-ctx.Done():
		return
	default:
		apiKey := "AIzaSyBOypms8DmfMpEx6-IRJzwz7lvBmE4kr94"
		for _, location := range data.Locations {
			var res GeoResponse
			adress := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
			geoWg.Add(1)
			go func(location string) {
				defer geoWg.Done()
				FetchGoRoutine(ctx, errChan, adress, &res, &geoWg, "locations")
				if res.Status != "OK" {
					errChan <- errors.New(res.Status)
					return
				}
				mutex.Lock()
				data.LocationsCoordinates[location] = fmt.Sprintf("%g,%g", res.Results[0].Geometry.Location.Lat, res.Results[0].Geometry.Location.Lng)
				mutex.Unlock()
			}(location)

		}
		geoWg.Wait()
	}
}
