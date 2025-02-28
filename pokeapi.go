package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/NHemmerly/poke-cli/internal/pokecache"
)

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  Pokedex
}

func getResponse(url string, cfg *config) ([]byte, error) {
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
	return body, nil
}

func PrintMapInfo(url string, cfg *config) (*LocationAreaNoID, error) {
	body, err := getResponse(url, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve body: %w", err)
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

func PrintLocationPokemon(url string, cfg *config) error {
	body, err := getResponse(url, cfg)
	if err != nil {
		return fmt.Errorf("could not retrieve body: %w", err)
	}
	var response Location
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Location not found")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func PrintPokeCatch(url string, cfg *config) error {
	body, err := getResponse(url, cfg)
	if err != nil {
		return fmt.Errorf("could not retrieve body: %w", err)
	}

	var response Pokemon
	if err = json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("pokemon name not found")
	}
	catchRate := float64(response.BaseExperience-36) / (650 - 36) // Make constants for these?
	if rand.Intn(100) >= int(catchRate*100) {
		fmt.Printf("%s was caught!\n", response.Name)
		cfg.pokedex.Add(&response)
	} else {
		fmt.Printf("%s escaped!\n", response.Name)
	}
	return nil
}
