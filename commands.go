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
	callback    func(cfg *config, arg string) error
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
		"explore": {
			name:        "explore",
			description: "Displays the names of pokemon at a map location: explore [location]",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches the selects pokemon: catch [pokemon name]",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints information about a pokemon if it is in the pokedex: inspect [pokemon name]",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints all pokemon caught by the current user",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Cannot close program")
}

func commandHelp(cfg *config, arg string) error {
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

func commandMap(cfg *config, arg string) error {
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

func commandMapb(cfg *config, arg string) error {
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

func commandExplore(cfg *config, location string) error {
	locationUrl := "https://pokeapi.co/api/v2/location-area/" + location
	fmt.Printf("Exploring %s...\n", location)
	err := PrintLocationPokemon(locationUrl, cfg)
	if err != nil {
		return fmt.Errorf("could not print location info: %w", err)
	}
	return nil
}

func commandCatch(cfg *config, name string) error {
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + name
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if err := PrintPokeCatch(pokemonUrl, cfg); err != nil {
		return fmt.Errorf("could not catch pokemon: %w", err)
	}
	return nil
}

func commandInspect(cfg *config, name string) error {
	val, ok := cfg.pokedex.Get(name)
	if !ok {
		fmt.Printf("You have not caught that pokemon\n")
		return nil
	}
	PrintPokemon(val)

	return nil
}

func commandPokedex(cfg *config, nul string) error {
	sortedPokemon := slices.Sorted(maps.Keys(cfg.pokedex.Pokemon))
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range sortedPokemon {
		fmt.Printf("\t- %s\n", pokemon)
	}
	return nil
}
