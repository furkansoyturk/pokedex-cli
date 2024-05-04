package pokeapi

type PokemonStruct struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
	BaseExperience int `json:"base_experience"`
}
