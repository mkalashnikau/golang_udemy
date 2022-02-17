package main

import (
	"log"

	"github.com/mkalashnikau/golang_udemy/helpers"
)

const numPool = 1000

// Channels are useful when our program has multiple packages

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	// We say: close the channel once it executed
	defer close(intChan)

	// Goroutine (A goroutine is a lightweight thread managed by the Go runtime.)
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
