package main

import (
	"fmt"
)

func commandHelp(cfg *config, input []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}

	fmt.Printf("\n")
	return nil
}
