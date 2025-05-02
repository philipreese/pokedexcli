package main

import "fmt"

func commandPokedex(config *cliConfig, args ...string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("No Pokemon in the Pokedex!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}