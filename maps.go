package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != "" {
		url = config.Next
	}
	return helperMap(config, url)
}

func commandMapb(config *Config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.Previous
	return helperMap(config, url)
}

func helperMap(config *Config, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("Failed status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var pArea pokeArea
	err = json.Unmarshal(body, &pArea)
	if err != nil {
		return err
	}

	config.Next = pArea.Next
	config.Previous = pArea.Previous

	for _, area := range pArea.Results {
		fmt.Println(area.Name)
	}

	return nil
}
