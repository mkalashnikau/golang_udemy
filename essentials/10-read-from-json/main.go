package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Each element corresponds to json structure
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
 [
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
 ]`
	var unmarshalled []Person

	// The function json.Unmarshal gets slice of bytes from a data source (myJson variable) and put the content to unmarshalled variable
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled) // output: [{Clark Kent black true} {Bruce Wayne black false}]

	// Write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Print(string(newJson))
}
