package main

import "strings"

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
