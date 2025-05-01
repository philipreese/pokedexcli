package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationAreas LocationAreas

	val, exists := cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, err
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&locationAreas)
		if err != nil {
			return LocationAreas{}, err
		}

		val, err = json.Marshal(&locationAreas)
		if err == nil {
			cache.Add(url, val)
		}
	}

	err := json.Unmarshal(val, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	return locationAreas, nil
}