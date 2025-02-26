package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = cleanInput(input)[0]
		if val, ok := Commands[input]; ok {
			err := val.callback()
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
