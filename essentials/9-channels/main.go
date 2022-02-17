package main

import (
	"log"
	"github.com/mkalashnikau/golang_udemy/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := 
	intChan <- helpers.RarandomNumber(numPool)
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
