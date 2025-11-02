package main

import (
	"fmt"
	"github.com/juandrzej/pokedex/internal/pokeapi"
	"math/rand"
	"os"
)

type Config struct {
	Next     string
	Previous string
	Client   *pokeapi.Client
	Pokedex  map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

var commands map[string]cliCommand

func commandExit(*Config, []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config, []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandMap(cfg *Config, args []string) error {
	list, err := cfg.Client.FetchLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	return commandMapHelper(cfg, list)
}

func commandMapb(cfg *Config, args []string) error {
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

func commandExplore(cfg *Config, args []string) error {
	list, err := cfg.Client.FetchLocationPokemons(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range list.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *Config, args []string) error {
	pokemonName := args[0]
	pokemon, err := cfg.Client.FetchPokemonData(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseExp := pokemon.BaseExperience
	userRoll := rand.Intn(300)
	if baseExp > userRoll {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		cfg.Pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	}
	fmt.Println(baseExp, userRoll)
	for name := range cfg.Pokedex {
		fmt.Println(name)
	}
	return nil
}

func commandInspect(cfg *Config, args []string) error {
	pokemonName := args[0]
	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("Pokemon not in pokedex yet.")
		return nil
	}
	fmt.Println("Name: " + pokemonName)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typePoke := range pokemon.Types {
		fmt.Println("  - " + typePoke.Type.Name)
	}
	return nil
}
