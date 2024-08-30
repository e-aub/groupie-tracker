package global

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func GetLocationsId(ctx context.Context, data *ArtistLocation, errChan chan error) {
	var geoWg sync.WaitGroup
	data.LocationsCoordinates = map[string]string{}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			apiKey := "AIzaSyBOypms8DmfMpEx6-IRJzwz7lvBmE4kr94"
			for _, location := range data.Locations {
				var res GeoResponse
				adress := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
				geoWg.Add(1)
				go func() {
					FetchGoRoutine(ctx, errChan, adress, &res, &geoWg, "")
					if res.Status != "OK" {
						errChan <- errors.New(res.Status)
						return
					}
					data.LocationsCoordinates[location] = fmt.Sprintf("%g,%g", res.Results[0].Geometry.Location.Lat, res.Results[0].Geometry.Location.Lng)
				}()

			}
			geoWg.Wait()
			return
		}
	}
}
