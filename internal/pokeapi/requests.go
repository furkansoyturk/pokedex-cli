package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationArea(url *string) LocationStruct {
	requestUrl := "https://pokeapi.co/api/v2/location-area/"

	if url != nil {
		requestUrl = *url
	}
	if val, ok := c.cache.Get(requestUrl); ok {
		return unmarshallResponse[LocationStruct](val)
	}
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(requestUrl, body)

	return unmarshallResponse[LocationStruct](body)
}

func (c *Client) ExplorePokemons(name string) PokemonStruct {
	requestUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(requestUrl, body)
	return unmarshallResponse[PokemonStruct](body)
}

func unmarshallResponse[T LocationStruct | PokemonStruct](body []byte) T {
	var response T
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
