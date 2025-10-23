package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

func FetchLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = baseUrl
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaList{}, fmt.Errorf("Failed status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	var areas LocationAreaList
	if err := json.Unmarshal(body, &areas); err != nil {
		return LocationAreaList{}, err
	}

	return areas, nil
}
