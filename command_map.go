package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, input []string) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = locationResp.Next
	cfg.previousLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, input []string) error {

	if cfg.previousLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.previousLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
