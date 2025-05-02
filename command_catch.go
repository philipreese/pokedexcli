package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/philipreese/pokedexcli/internal/pokeapi"
)

func commandCatch(config *cliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	minExp := 36
	maxExp := 608
	minProb := 0.7
	maxProb := 0.05
	interpolationFactor := float64(pokemon.BaseExperience - minExp) / float64(maxExp - minExp)
	catchProbability := minProb + (maxProb - minProb)*interpolationFactor
	
	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}