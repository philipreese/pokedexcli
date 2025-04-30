package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandMap = map[string]cliCommand {}

func main() {
	commandMap = map[string]cliCommand {
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
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, ok := commandMap[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
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