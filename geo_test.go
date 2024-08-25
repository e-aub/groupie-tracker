package main

// import (
// 	"fmt"
// 	"testing"
// )

// func TestGeo(t *testing.T) {
// 	apiKey := "AIzaSyBOypms8DmfMpEx6-IRJzwz7lvBmE4kr94"
// 	var locations Index
// 	err := Fetch("https://groupietrackers.herokuapp.com/api/locations", &locations)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Error(err)
// 	}
// 	for _, locationsGroup := range locations.Location1 {
// 		for _, location := range locationsGroup.Location {
// 			var res GeoResponse
// 			adress := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)
// 			err := Fetch(adress, &res)
// 			if err != nil {
// 				t.Error(err)
// 				continue
// 			}
// 			if res.Status != "OK" {
// 				t.Error("eroooooor")
// 				continue
// 			}

// 		}
// 	}
// }
