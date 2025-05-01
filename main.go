package main

import "github.com/philipreese/pokedexcli/internal/pokeapi"

func main() {
	config := &cliConfig{
		pokedex: map[string]pokeapi.Pokemon{},
	}

	startRepl(config)
}