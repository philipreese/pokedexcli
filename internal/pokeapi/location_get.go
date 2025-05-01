package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	var locationArea LocationArea

	val, exists := cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		val, err = json.Marshal(&locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		cache.Add(url, val)
	}

	err := json.Unmarshal(val, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}