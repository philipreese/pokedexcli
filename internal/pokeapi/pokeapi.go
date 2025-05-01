package pokeapi

import (
	"github.com/philipreese/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2" 

var cache = pokecache.NewCache(5)
