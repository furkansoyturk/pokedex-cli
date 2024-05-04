package main

import (
	"github.com/furkansoyturk/pokedex-cli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeClient:    pokeClient,
		caughtPokeman: map[string]pokeapi.PokemonStruct{},
	}

	StartRepl(cfg)
}
