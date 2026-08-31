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

func commandCall(cfg *config, cleanedInput []string) {
	commandName := cleanedInput[0]
	args := cleanedInput[1:]

	cmd, ok := cfg.commands[commandName]
	if !ok {
		fmt.Println("Unknown command")
		return
	}

	if err := cmd.callback(cfg, args...); err != nil {
		fmt.Println(err)
	}
}
