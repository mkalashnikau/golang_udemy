package main

import "log"

func main() {
	myNum := 100
	isTrue := false

	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and true")
	} else if myNum < 100 && isTrue {

	} else if myNum == 1-1 || isTrue {

	} else if myNum > 100 && isTrue == false {
		log.Panicln("3'")
	}

	myVar := "brra"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")

	case "dog":
		log.Println("Cat is set to dog")

	default:
		log.Println("cat is something else")
	}

}
