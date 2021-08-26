package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	MAX_REQUESTS = 800
)

func main() {
	numberOfRequests := readNumberOfRequestsFromCL()

	url := "https://pokeapi.co/api/v2/pokemon/%d"

	for i := 1; i <= numberOfRequests; i++ {
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

func readNumberOfRequestsFromCL() int {
	args := os.Args[1:]
	numberOfRequests := MAX_REQUESTS
	if len(args) >= 1 {
		var err error
		numberOfRequests, err = strconv.Atoi(args[0])
		if err != nil {
			numberOfRequests = MAX_REQUESTS
		}
	}

	return numberOfRequests
}
