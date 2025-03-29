package main

import "os"

func commandExit(cfg *config) error {
	print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
