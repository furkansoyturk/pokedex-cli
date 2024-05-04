package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	name := args[0]
	response := cfg.pokeClient.CatchPokemons(name)
	catchingProb := response.BaseExperience
	fmt.Printf("Throwing a pokeball at %s", name)
	random := rand.Intn(catchingProb)
	fmt.Printf("\n math random : %v", random)
	if random > 50 {
		fmt.Printf("\n %s was caught", name)
	} else {
		fmt.Printf("\n %s escaped", name)
	}

	fmt.Println("")
	return nil
}
