package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/juandrzej/pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

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

func (c *Client) FetchLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = baseUrl
	}

	if b, ok := c.cache.Get(url); ok {
		var areas LocationAreaList
		if err := json.Unmarshal(b, &areas); err != nil {
			return LocationAreaList{}, err
		}
		return areas, nil
	}

	res, err := c.httpClient.Get(url)
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

	c.cache.Add(url, body)

	var areas LocationAreaList
	if err := json.Unmarshal(body, &areas); err != nil {
		return LocationAreaList{}, err
	}
	return areas, nil
}
