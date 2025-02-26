package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
