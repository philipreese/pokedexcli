package main

import (
	"errors"
	"fmt"

	"github.com/philipreese/pokedexcli/internal/pokeapi"
)

func commandExplore(config *cliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]

	locationArea, err := pokeapi.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}