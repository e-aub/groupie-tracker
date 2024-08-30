package global

type (
	Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
	}
	Locations struct {
		Index []ArtistLocation `json:"index"`
	}
	ArtistLocation struct {
		Id                   int      `json:"id"`
		Locations            []string `json:"locations"`
		LocationsCoordinates map[string]string
	}
	ArtistDate struct {
		Id   int      `json:"id"`
		Date []string `json:"dates"`
	}

	ArtistRelation struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
	GeoResponse struct {
		Results []struct {
			PlaceID  string `json:"place_id"`
			Geometry struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
		} `json:"results"`
		Status string `json:"status"`
	}
)
