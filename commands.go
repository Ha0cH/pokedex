package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getPokemonsInArea(cfg *config, area string) ([]string, error) {
	prefix := "https://pokeapi.co/api/v2/location-area/"
	url := prefix + area

	var data []byte
	data, ok := cfg.cache.Get(url)
	if !ok {
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		cfg.cache.Add(url, data)
	}

	var areaDetail LocationAreaDetail
	err := json.Unmarshal(data, &areaDetail)
	if err != nil {
		return nil, err
	}

	pokemons := areaDetail.PokemonEncounters
	var listPokemons []string
	for _, p := range pokemons {
		listPokemons = append(listPokemons, p.Pokemon.Name)
	}

	return listPokemons, nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Invalid input: 'explore' must contain an area name")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemons:")

	areaName := args[0]
	pokemons, err := getPokemonsInArea(cfg, areaName)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for c, s := range cfg.commands {
		fmt.Printf("\n%s: %s", c, s.description)
	}
	fmt.Print("\n")

	return nil
}

func getLocationAreas(cfg *config, url string) ([]string, error) {
	var data []byte

	data, ok := cfg.cache.Get(url)
	if !ok {
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
		}

		// Read body into byte
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		//cache it
		cfg.cache.Add(url, data)
	}

	// Whether cached or fresh, decode the same []byte
	var page LocationAreasPage
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, err
	}

	// update URLs in the config
	cfg.nextURL = page.Next
	cfg.previousURL = page.Previous

	var areas []string
	for _, a := range page.Results {
		areas = append(areas, a.Name)
	}

	return areas, nil
}

func commandMap(cfg *config, args ...string) error {
	areas, err := getLocationAreas(cfg, cfg.nextURL)
	if err != nil {
		return err
	}

	for _, item := range areas {
		fmt.Println(item)
	}

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.previousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := getLocationAreas(cfg, cfg.previousURL)
	if err != nil {
		return err
	}

	for _, item := range areas {
		fmt.Println(item)
	}

	return nil
}

func getPokemon(cfg *config, name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	data, ok := cfg.cache.Get(url)
	if !ok {
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("request failed with status code %d", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		cfg.cache.Add(url, data)
	}

	var p Pokemon
	err := json.Unmarshal(data, &p)
	if err != nil {
		return Pokemon{}, err
	}

	return p, nil

}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Input must contain a valid pokemon name")
	}

	pokemon, err := getPokemon(cfg, args[0])
	if err != nil {
		return err
	}

	_, ok := cfg.pokedex[pokemon.Name]
	if ok {
		fmt.Println("you already caught this pokemon")
		return nil
	}

	// try to catch it
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchChance := 100 - pokemon.BaseExperience/3
	if catchChance < 5 {
		catchChance = 5
	}

	roll := rand.Intn(100) + 1

	if roll <= catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
