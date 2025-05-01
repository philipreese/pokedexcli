package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliConfig struct {
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(config *cliConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(config); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapBack,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}