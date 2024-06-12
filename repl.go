package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		"map": {
			name:        "map",
			description: "Get the next map page",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous map page",
			// callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
	}
}

func correctInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
func repl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		command := correctInput(scanner.Text())
		if len(command) == 0 {
			continue
		}

		commandName := command[0]
		cmd, ok := getCommands()[commandName]

		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command, sorry!")
			continue
		}
	}
}
