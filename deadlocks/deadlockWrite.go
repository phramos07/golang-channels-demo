package main

import "fmt"

func main() {
	intChan := make(chan int)

	intChan <- 2

	fmt.Printf("Unreachable code.")
}
