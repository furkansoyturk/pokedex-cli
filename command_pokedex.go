package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	m := cfg.caughtPokeman
	if len(m) == 0 {
		fmt.Println("Pokedex has no pokemon")
	} else {
		fmt.Println("Your pokedex: ")
		for _, value := range m {
			fmt.Printf(" - %v", value.Name)
			fmt.Println("")
		}
	}
	return nil
}
