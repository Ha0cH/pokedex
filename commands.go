package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
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

func commandMap(cfg *config) error {
	areas, err := getLocationAreas(cfg, cfg.nextURL)
	if err != nil {
		return err
	}

	for _, item := range areas {
		fmt.Println(item)
	}

	return nil
}

func commandMapBack(cfg *config) error {
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
