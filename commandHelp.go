package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Oh dear! Lets get you some help...")
	for _, cmd := range getCommands() {
		fmt.Printf("available commands: %s, %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
