package main

import (
	"fmt"
	"os"

	pokeapi "github.com/philipreese/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
}

type cliConfig struct {
	next  	 *string
	previous *string
}

func commandExit(config *cliConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *cliConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range(commandMenu) {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *cliConfig) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.next != nil {
		url = *config.next
	}

	return handleLocationArea(url)
}

func commandMapBack(config *cliConfig) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area"
	if config.previous != nil {
		url = *config.previous
	}
	
	return handleLocationArea(url)
}

func handleLocationArea(url string) error {
	locationAreas, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	config.next = locationAreas.Next
	config.previous = locationAreas.Previous

	for _, locationArea := range(locationAreas.Results) {
		fmt.Println(locationArea.Name)	
	}
	return nil
}