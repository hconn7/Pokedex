package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// Cli Stuct

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// CliCommands

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "I cannot help you right now",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
	}
}
