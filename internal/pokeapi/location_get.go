package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	var location Location

	val, exists := cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&location)
		if err != nil {
			return Location{}, err
		}

		val, err = json.Marshal(&location)
		if err != nil {
			return Location{}, err
		}

		cache.Add(url, val)
	}

	err := json.Unmarshal(val, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}