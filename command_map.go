package main

import (
	"fmt"
	"github.com/furkansoyturk/pokedex-cli/internal/pokeapi"
)

func commandMap() error {
	response := pokeapi.GetLocationArea()
	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	return nil
}
