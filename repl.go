package main

import (
	"bufio"
	"fmt"
	"github.com/abdielrumaldo/pokedex/internal/pokeapi"
	"github.com/abdielrumaldo/pokedex/internal/pokecache"
	"os"
	"strings"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
	cache               pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, input []string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())

		command, ok := getCommands()[input[0]]
		if ok {
			err := command.callback(cfg, input)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown Command\n")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location areas in order",
			callback:    commandMapf,
		},

		"pmap": {
			name:        "pmap",
			description: "Displays previous location areas in order",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "Displays the pokemon that can be found in provided area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Throw a pokemon at the selected pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect the attributes of the selected pokemon if caught.",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Return all pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}

}
