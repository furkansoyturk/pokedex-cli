package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocationArea(url *string) LocationStruct {
	requestUrl := "https://pokeapi.co/api/v2/location-area/"

	if url != nil {
		requestUrl = *url
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

	return unmarshallLocation(body)
}

func unmarshallLocation(body []byte) LocationStruct {
	locationStruct := LocationStruct{}

	err := json.Unmarshal(body, &locationStruct)

	if err != nil {
		fmt.Println(err)
	}

	return locationStruct
}
