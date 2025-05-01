package pokeapi

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource 		  `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterVersionDetails struct {
	Rate    int     		 `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string   		  `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource 		`json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter 	  `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int    			   `json:"min_level"`
	MaxLevel        int    			   `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int    			   `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}