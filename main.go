package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandMenu = map[string]cliCommand {}
var config = cliConfig{next: nil, previous: nil}

func main() {
	commandMenu = map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapBack,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, ok := commandMenu[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(&config); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}