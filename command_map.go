package main

import (
	"errors"
	"fmt"
	"github.com/furkansoyturk/pokedex-cli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	response := pokeapi.GetLocationArea(cfg.nextLocURL)
	cfg.nextLocURL = response.Next
	cfg.prevLocURL = response.Previous

	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocURL == nil {
		return errors.New("No previous location !!!")
	}
	response := pokeapi.GetLocationArea(cfg.prevLocURL)

	cfg.nextLocURL = response.Next
	cfg.prevLocURL = response.Previous

	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	return nil
}
