package main

import (
	"fmt"
)

func commandExplore(cfg *config, input []string) error {

	if len(input) < 2 {
		return fmt.Errorf("No area provided in command '%v'\n", input)
	}

	locationResp, err := cfg.pokeapiClient.GetLocation(input[1])
	if err != nil {
		return err
	}

	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
