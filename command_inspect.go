package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	name := args[0]
	pokemon, ok := cfg.caughtPokeman[name]
	if ok {
		response := cfg.pokeClient.CatchPokemons(pokemon.Name)
		fmt.Printf("\n name: %s", response.Name)
		fmt.Printf("\n height: %v ", response.Height)
		fmt.Printf("\n base experience: %v", response.BaseExperience)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Println("")
	return nil
}
