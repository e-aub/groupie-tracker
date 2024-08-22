package global

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseUrl string = "https://groupietrackers.herokapp.com/api"

// https://groupietrackers.herokuapp.com/api

func Read(w http.ResponseWriter, err *error, url string, data any, wg *CheckWG) {
	if !wg.NotWG {
		defer wg.WG.Done()
	}
	res, errr := http.Get(BaseUrl + url)
	if errr != nil {
		// http.Error(w, "page not found", http.StatusNotFound)
		fmt.Println("khhh")
		err = &errr
		return
		

	}
	fmt.Println("khhh")
	defer res.Body.Close()
	errr = json.NewDecoder(res.Body).Decode(data)
	if errr != nil {
		// http.Error(w, "page not found", http.StatusNotFound)
		err = &errr
		return
	}
}
