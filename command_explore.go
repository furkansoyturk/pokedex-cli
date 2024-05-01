package main

import (
	"fmt"
)

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
