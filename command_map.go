package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	response := cfg.pokeClient.GetLocationArea(cfg.nextLocURL)
	cfg.nextLocURL = response.Next
	cfg.prevLocURL = response.Previous

	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("No previous location !!!")
	}
	response := cfg.pokeClient.GetLocationArea(cfg.prevLocURL)

	cfg.nextLocURL = response.Next
	cfg.prevLocURL = response.Previous

	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	return nil
}
func commandExplore(cfg *config, args ...string) error {
	name := args[0]
	response := cfg.pokeClient.ExplorePokemons(name)
	fmt.Println("Following pokemons found:")
	for _, pokemons := range response.PokemonEncounters {
		fmt.Printf("%s\n", pokemons.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}
