package main

import "fmt"

func commandHelp() error {
	fmt.Println("Oh dear! Lets get you some help...")
	for _, cmd := range getCommands() {
		fmt.Println("available commands: %s, %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
