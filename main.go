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
