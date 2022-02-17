package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Map is [index]value
	myMap := make(map[string]User)

	me := User{
		FirstName: "Mikita",
		LastName:  "Kalashnikau",
	}
	myMap["me"] = me
	log.Println(myMap["me"].FirstName)

	// Declare slice
	var mySlice []string
	mySlice = append(mySlice, "Mikita")
	mySlice = append(mySlice, "Stuart")
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	// Print 1-3 numbers
	log.Println(numbers[1:3])
}
