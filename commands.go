package main

import (
	"fmt"
	"github.com/juandrzej/pokedex/internal/pokeapi"
	"os"
)

type Config struct {
	Next     string
	Previous string
	Client   *pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var commands map[string]cliCommand

func commandExit(*Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandMap(cfg *Config) error {
	list, err := cfg.Client.FetchLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	return commandMapHelper(cfg, list)
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	list, err := cfg.Client.FetchLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	return commandMapHelper(cfg, list)
}

func commandMapHelper(cfg *Config, list pokeapi.LocationAreaList) error {
	cfg.Next = list.Next
	cfg.Previous = list.Previous
	for _, area := range list.Results {
		fmt.Println(area.Name)
	}
	return nil
}
