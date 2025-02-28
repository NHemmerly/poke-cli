package main

import (
	"errors"
	"fmt"
	"maps"
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

	response, err := PrintMapInfo(cfg.next, cfg)
	if err != nil {
		return fmt.Errorf("could not print map: %w", err)
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

	response, err := PrintMapInfo(cfg.previous, cfg)
	if err != nil {
		return fmt.Errorf("could not print map: %w", err)
	}

	cfg.next = response.Next
	cfg.previous = response.Previous

	return nil
}
