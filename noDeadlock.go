package main

import "fmt"

func main() {
	intChan := make(chan int)

	go func() {
		rec := <-intChan // happens after line 13
		fmt.Printf("Received value: %d\n", rec)
	}()

	intChan <- 2 // happens before line 9
}
