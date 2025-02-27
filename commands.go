package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os"
	"slices"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

var Commands map[string]cliCommand

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations, each subsequent call displays the next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations of the map command",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Cannot close program")
}

func commandHelp(cfg *config) error {
	descriptions := func() string {
		var out string
		sortedKeys := slices.Sorted(maps.Keys(Commands))
		for _, key := range sortedKeys {
			out += fmt.Sprintf("%s: %s\n", key, Commands[key].description)
		}
		return out
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:
	
%s`, descriptions())
	return nil
}

func commandMap(cfg *config) error {
	if cfg.next == "" {
		cfg.next = "https://pokeapi.co/api/v2/location-area"
	}

	cacheData, found := cfg.cache.Get(cfg.next)

	var body []byte

	if found {
		fmt.Println("Using cached data...")
		body = cacheData
	} else {
		res, err := http.Get(cfg.next)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %w", err)
		}

		cfg.cache.Add(cfg.next, body)
	}

	var response LocationAreaNoID
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("could not unmarshal: %w", err)
	}

	for _, name := range response.Results {
		fmt.Println(name.Name)
	}

	cfg.next = response.Next
	cfg.previous = response.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var body []byte
	cacheData, found := cfg.cache.Get(cfg.previous)
	if found {
		fmt.Println("Using cached data...")
		body = cacheData
	} else {
		res, err := http.Get(cfg.previous)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %w", err)
		}
		cfg.cache.Add(cfg.previous, body)
	}

	var response LocationAreaNoID
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("could not unmarshal: %w", err)
	}

	for _, name := range response.Results {
		fmt.Println(name.Name)
	}

	cfg.next = response.Next
	cfg.previous = response.Previous

	return nil
}
