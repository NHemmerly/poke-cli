package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	cfg := &config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = cleanInput(input)[0]
		if val, ok := Commands[input]; ok {
			err := val.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
