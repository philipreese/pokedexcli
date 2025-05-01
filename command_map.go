package main

import (
	"errors"
	"fmt"

	"github.com/philipreese/pokedexcli/internal/pokeapi"
)

func commandMap(config *cliConfig) error {
	return handleLocationArea(config, true)
}

func commandMapBack(config *cliConfig) error {
	if config.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	
	return handleLocationArea(config, false)
}

func handleLocationArea(config *cliConfig, next bool) error {
	var url *string
	if next {
		url = config.nextLocationsURL
	} else {
		url = config.previousLocationsURL
	}

	locationAreas, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationAreas.Next
	config.previousLocationsURL = locationAreas.Previous

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)	
	}
	return nil
}