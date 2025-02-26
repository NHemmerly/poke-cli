package main

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return []string{}
	}
	out := strings.Fields(text)
	for i, word := range out {
		word = strings.TrimSpace(word)
		out[i] = strings.ToLower(word)
	}
	return out
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Cannot close program")
}

func commandHelp() error {
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
