package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

}

func cleanInput(text string) []string {
	afterSplit := strings.Fields(text)
	for i, char := range afterSplit {
		afterSplit[i] = strings.ToLower(char)
	}
	return afterSplit
}
