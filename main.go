package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			fmt.Println("Cannot process empty command, please write sth.")
			continue
		}
		splitText := cleanInput(text)
		fmt.Printf("Your command was: %s\n", splitText[0])
	}

}

func cleanInput(text string) []string {
	afterSplit := strings.Fields(text)
	for i, char := range afterSplit {
		afterSplit[i] = strings.ToLower(char)
	}
	return afterSplit
}
