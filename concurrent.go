package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	MAX_REQUESTS = 800
	URL          = "https://pokeapi.co/api/v2/pokemon/%d"
)

func main() {
	numberOfRequests := readNumberOfRequestsFromCL()

	wg := sync.WaitGroup{}
	pokeChan := make(chan map[string]interface{})
	endChan := make(chan bool)
	errChan := make(chan error)

	for i := 1; i <= numberOfRequests; i++ {
		pokeUrl := fmt.Sprintf(URL, i)

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			res, err := http.Get(pokeUrl)
			if err != nil {
				errChan <- err
				return
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				errChan <- err
				return
			}
			defer res.Body.Close()

			var pokemon map[string]interface{}
			err = json.Unmarshal(body, &pokemon)
			if err != nil {
				errChan <- err
				return
			}

			pokeChan <- pokemon
		}(pokeUrl)
	}

	go func() {
		wg.Wait()
		endChan <- true
	}()

	for {
		select {
		case <-endChan:
			return
		case poke := <-pokeChan:
			fmt.Printf("Name: %s, Height: %f, Weight: %f\n", poke["name"], poke["height"], poke["weight"])
		case err := <-errChan:
			fmt.Printf("Error during request: %s\n", err.Error())
		}
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
