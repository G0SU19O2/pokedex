package main

import (
	"fmt"
)

func commandMap(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationURL)
	if err != nil {
		return err
	}
	config.nextLocationURL = locationsResp.Next
	config.prevLocationURL = locationsResp.Previous
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.prevLocationURL)
	if err != nil {
		return err
	}
	config.nextLocationURL = locationsResp.Next
	config.prevLocationURL = locationsResp.Previous
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
