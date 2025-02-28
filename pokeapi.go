package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NHemmerly/poke-cli/internal/pokecache"
)

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
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

func PrintMapInfo(url string, cfg *config) (*LocationAreaNoID, error) {
	var body []byte
	cacheData, found := cfg.cache.Get(url)
	if found {
		fmt.Println("Using cached data...")
		body = cacheData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body: %w", err)
		}
		cfg.cache.Add(url, body)
	}

	var response LocationAreaNoID
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("could not unmarshal: %w", err)
	}

	for _, name := range response.Results {
		fmt.Println(name.Name)
	}
	return &response, nil
}
