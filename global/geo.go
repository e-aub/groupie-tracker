package global

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
		for _, location := range data.Locations {
			var response GeoResponse
			geoWg.Add(1)
			go func(location string) {
				defer geoWg.Done()
				select {
				case <-ctx.Done():
					return
				default:
					url := fmt.Sprintf("https://maps-data.p.rapidapi.com/geocoding.php?query=%s&lang=en&country=fr", location)
					req, err := http.NewRequest("GET", url, nil)
					if err != nil {
						errChan <- err
						return
					}

					req.Header.Add("x-rapidapi-key", "3deb8955ecmsh5bfc47a8a9587b3p19cb4ajsn0a1824e4a63c")
					req.Header.Add("x-rapidapi-host", "maps-data.p.rapidapi.com")

					res, err := http.DefaultClient.Do(req)
					if res.StatusCode != 200 {
						errChan <- fmt.Errorf("status not 200 : %d", res.StatusCode)
						return
					} else if err != nil {
						errChan <- err
						return
					}
					defer res.Body.Close()

					err = json.NewDecoder(res.Body).Decode(&response)
					if err != nil {
						errChan <- err
						return
					}
				}
				mutex.Lock()
				data.LocationsCoordinates[location] = fmt.Sprintf("%g,%g", response.Location.Lat, response.Location.Lng)
				defer mutex.Unlock()
			}(location)

		}
		geoWg.Wait()
	}
}
