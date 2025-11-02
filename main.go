package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/juandrzej/pokedex/internal/pokeapi"
)

func main() {

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "It displays the names of  (next) 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "It displays the names of (previous) 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "It displays a list of all the Pokémon located in a given location.",
			callback:    commandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{Client: pokeapi.NewClient(5 * time.Second)}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(config, input[1:]); err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(text string) []string {
	afterSplit := strings.Fields(text)
	for i, char := range afterSplit {
		afterSplit[i] = strings.ToLower(char)
	}
	return afterSplit
}
