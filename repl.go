package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/G0SU19O2/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func startRepl(cfg *config) {
	supportedCommands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command, ok := supportedCommands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println("Something wrong")
		}
	}
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
			description: "Display list of location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display list of previous location",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Exploring",
			callback:    commandExplore,
		},
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
