package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Ha0cH/pokedex/internal/pokecache"
)

func main() {
	cfg := &config{
		cache: pokecache.NewCache(time.Minute * 5),
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Display a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Display the next 20 location areas in the Pokemon world",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Display the previous 20 location areas in the Pokemon world",
				callback:    commandMapBack,
			},
		},
		nextURL:     "https://pokeapi.co/api/v2/location-area/",
		previousURL: "",
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			break
		}

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandCall(cfg, cleaned[0])
	}

}
