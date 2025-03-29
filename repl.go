package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	page int
}

func (c *config) Next() string {
	c.page++
	return fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=20&offset=%d", c.page*20)
}
func (c *config) Previous() string {
	if c.page > 0 {
		c.page--
	}
	return fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=20&offset=%d", c.page*20)
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
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
		err := command.callback(cfg)
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
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
