package main

import "fmt"

func commandPokedex(cfg *config, input []string) error {

	if len(cfg.pokeapiClient.Pokedex) < 1 {
		fmt.Println("you haven't caught any pokemon")
		return nil
	}
	for _, poke := range cfg.pokeapiClient.Pokedex {
		fmt.Printf(" - %s \n", poke.Name)
	}
	return nil
}
