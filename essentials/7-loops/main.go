package main

import "log"

func main() {

	// Example 1
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	// Example 2
	mapAnimals := make(map[string]string)
	mapAnimals["dog"] = "Fide"
	mapAnimals["cat"] = "Fluffy"

	for animalType, animal := range mapAnimals {
		log.Println(animalType, animal)
	}

	// Example 3
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@gmaol.com", 27})
	users = append(users, User{"Borjomi", "Young", "borjomi@gmaol.com", 30})
	users = append(users, User{"Sally", "Russell", "sally@gmaol.com", 25})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age)
	}
}
