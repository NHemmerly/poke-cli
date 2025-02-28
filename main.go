package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/NHemmerly/poke-cli/internal/pokecache"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	cfg := &config{cache: pokecache.NewCache(20 * time.Second), pokedex: NewPokedex()}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputList := cleanInput(input)
		command, arg := "", ""
		if len(inputList) > 0 {
			command = inputList[0]
			if len(inputList) > 1 {
				arg = inputList[1]
			}
		}
		if val, ok := Commands[command]; ok {
			err := val.callback(cfg, arg)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
