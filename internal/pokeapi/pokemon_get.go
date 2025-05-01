package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var pokemon Pokemon

	val, exists := cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		val, err = json.Marshal(&pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, val)
	}

	err := json.Unmarshal(val, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}