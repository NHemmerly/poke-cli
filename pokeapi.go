package main

type config struct {
	next     string
	previous string
}

type (
	LocationAreaNoID struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []Name `json:"results"`
	}
	LocationResponse struct {
		Id                     int                   `json:"id"`
		Name                   string                `json:"name"`
		Game_index             int                   `json:"game_index"`
		Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
		Location               NamedAPIResource      `json:"location"`
		Names                  []Name                `json:"names"`
		Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
	}
	EncounterMethodRate struct {
		Encounter_method NamedAPIResource        `json:"encounter_method"`
		Version_details  EncounterVersionDetails `json:"version_details"`
	}
	EncounterVersionDetails struct {
		Rate    int              `json:"rate"`
		Version NamedAPIResource `json:"version"`
	}
	PokemonEncounter struct {
		Pokemon         NamedAPIResource        `json:"pokemon"`
		Version_details EncounterVersionDetails `json:"version_details"`
	}
	NamedAPIResource struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	Name struct {
		Name     string           `json:"name"`
		Language NamedAPIResource `json:"language"`
	}
)
