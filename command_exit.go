package main

import "os"

func commandExit() error {
	print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
