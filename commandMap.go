package main

import (
	"fmt"
	"log"
)
import pokeapi "github.com/hconn7/Pokedex.git/internal/pokeApi"

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)

	}
	return nil
}
