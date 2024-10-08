package global

type (
	Error struct {
		Message string
		Code    int
	}
	Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
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
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"data"`
	}
)
