package main

import "fmt"

func main() {
	intChan := make(chan int)

	intChan <- 2

	rec := <-intChan
	fmt.Printf("A value will never be read from the channel above, thus this will never be printed: %d", rec)
}
