package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocationArea() LocationStruct {

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
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
