package main

import (
	// "encoding/json"
	"fmt"
)

func commandInspect(cfg *config, input []string) error {

	// Check if we have seen this pokemon
	pokeData, ok := cfg.pokeapiClient.Pokedex[input[1]]
	//If we have not, give an error with a reason
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	// if yes, detail the following
	// pokeJson, err := json.Marshal(pokeData)
	// if err != nil {
	// 	fmt.Print("Error Marshaling pokeData")
	// 	return err
	// }
	// fmt.Print(string(pokeJson))
	fmt.Printf("Name: %s\nHeight: %d\n", pokeData.Name, pokeData.Height)

	fmt.Println("Stats:")
	for _, val := range pokeData.Stats {
		fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range pokeData.Types {
		fmt.Printf("  -%s\n", val.Type.Name)
	}
	/*
		Name: pidgey
		Height: 3
		Weight: 18
		Stats:
		  -hp: 40
		  -attack: 45
		  -defense: 40
		  -special-attack: 35
		  -special-defense: 35
		  -speed: 56
		Types:
		  - normal
		  - flying
	*/
	return nil
}
