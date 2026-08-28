package main

type config struct {
	commands    map[string]cliCommand
	nextURL     string
	previousURL string
}
