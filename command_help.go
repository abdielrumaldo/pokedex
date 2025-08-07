package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	return nil
}
