package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, input []string) error {

	if len(input) < 2 {
		return fmt.Errorf("No pokemon provided in command '%v'\n", input)
	}

	// Get pokemon information
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(input[1])
	if err != nil {
		return err
	}

	// Check for existence in pokedex
	_, ok := cfg.pokeapiClient.Pokedex[pokemonResp.Name]
	if ok {
		fmt.Printf("You already caught %s.\n", pokemonResp.Name)
		return nil
	}

	//Attempting catch
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	catchVal := rand.Intn(pokemonResp.BaseExperience)

	if catchVal > pokemonResp.BaseExperience/2 {
		fmt.Printf("%s was caught! %d\n", pokemonResp.Name, catchVal)
		cfg.pokeapiClient.Pokedex[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped! %d\n", pokemonResp.Name, catchVal)
	}

	return nil
}
