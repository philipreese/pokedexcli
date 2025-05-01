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
	
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}