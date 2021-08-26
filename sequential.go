package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	MAX_POKEMON = 800
)

func main() {
	url := "https://pokeapi.co/api/v2/pokemon/%d"

	for i := 1; i < MAX_POKEMON; i++ {
		pokeUrl := fmt.Sprintf(url, i)

		res, err := http.Get(pokeUrl)
		if err != nil {
			fmt.Printf("Couldn't complete the request: %s", err.Error())
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Couldn't read body from response: %s", err.Error())
			return
		}
		defer res.Body.Close()

		var pokemon map[string]interface{}
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Printf("Couldn't deserialize json into object: %s", err.Error())
			return
		}

		fmt.Printf("Name: %s, Height: %f, Weight: %f\n", pokemon["name"], pokemon["height"], pokemon["weight"])
	}
}
