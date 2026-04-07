package PokeAPI

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type LocationPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationPage(url string) LocationPage {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locationPage LocationPage
	err = json.Unmarshal(body, &locationPage)
	if err != nil {
		fmt.Println(err)
	}

	return locationPage
}

func (locationPage LocationPage) Print() {
	for _, place := range locationPage.Results {
		fmt.Println(place.Name)
	}
}


