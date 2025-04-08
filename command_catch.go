package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokemon at %s...\n", pokemonName)
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	if pokemon, ok := config.caughtPokemon[pokemon.Name]; ok {
		fmt.Printf("%s was caught\n", pokemon.Name)
		return nil
	}
	res := rand.IntN(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught\n", pokemon.Name)
	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
