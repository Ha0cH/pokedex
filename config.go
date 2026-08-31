package main

import "github.com/Ha0cH/pokedex/internal/pokecache"

type config struct {
	commands    map[string]cliCommand
	nextURL     string
	previousURL string
	cache       *pokecache.Cache
	pokedex     map[string]Pokemon
}
