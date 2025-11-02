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

const baseURL = "https://pokeapi.co/api/v2/location-area/"

func (c *Client) FetchLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = baseURL + "?offset=0&limit=20"
	}

	if b, ok := c.cache.Get(url); ok {
		fmt.Println("[cache] hit:", url)
		var areas LocationAreaList
		if err := json.Unmarshal(b, &areas); err != nil {
			return LocationAreaList{}, err
		}
		return areas, nil
	} else {
		fmt.Println("[cache] miss:", url)
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

type LocationAreaInfoList struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) FetchLocationPokemons(location string) (LocationAreaInfoList, error) {
	url := baseURL + location

	if b, ok := c.cache.Get(url); ok {
		fmt.Println("[cache] hit:", url)
		var infos LocationAreaInfoList
		if err := json.Unmarshal(b, &infos); err != nil {
			return LocationAreaInfoList{}, err
		}
		return infos, nil
	} else {
		fmt.Println("[cache] miss:", url)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaInfoList{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaInfoList{}, fmt.Errorf("Failed status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaInfoList{}, err
	}

	c.cache.Add(url, body)

	var infos LocationAreaInfoList
	if err := json.Unmarshal(body, &infos); err != nil {
		return LocationAreaInfoList{}, err
	}
	return infos, nil
}
