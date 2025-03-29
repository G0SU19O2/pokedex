package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(config *config) error {
	res, err := http.Get(config.Next())
	if err != nil {
		return err
	}
	var locations Locations
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}


func commandMapB(config *config) error {
	res, err := http.Get(config.Previous())
	if err != nil {
		return err
	}
	var locations Locations
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

type Locations struct {
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
