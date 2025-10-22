package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	}

	scanner := bufio.NewScanner(os.Stdin)

	config := Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		command, ok := commands[text]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(&config); err != nil {
			fmt.Println(err)
		}

		// splitText := cleanInput(text)
		// fmt.Printf("Your command was: %s\n", splitText[0])
	}

}

func cleanInput(text string) []string {
	afterSplit := strings.Fields(text)
	for i, char := range afterSplit {
		afterSplit[i] = strings.ToLower(char)
	}
	return afterSplit
}

func commandExit(config *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

var commands map[string]cliCommand

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}
