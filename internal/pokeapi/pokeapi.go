package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) FetchLocationAreas(url string) (LocationAreaList, error) {
	if url == "" {
		url = baseURLLocation + "?offset=0&limit=20"
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

func (c *Client) FetchLocationPokemons(location string) (LocationAreaInfoList, error) {
	url := baseURLLocation + location

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

func (c *Client) FetchPokemonData(pokemonName string) (Pokemon, error) {
	url := baseURLPokemon + pokemonName

	if b, ok := c.cache.Get(url); ok {
		fmt.Println("[cache] hit:", url)
		var pokemon Pokemon
		if err := json.Unmarshal(b, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	} else {
		fmt.Println("[cache] miss:", url)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Failed status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
