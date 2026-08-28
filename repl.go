package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)
	lowered := strings.ToLower(cleaned)
	words := strings.Split(lowered, " ")

	return words
}

func commandCall(cfg *config, cleanedInput string) {
	cmd, ok := cfg.commands[cleanedInput]
	if !ok {
		fmt.Println("Unknown command")
		return
	}

	if err := cmd.callback(cfg); err != nil {
		fmt.Println(err)
	}
}
