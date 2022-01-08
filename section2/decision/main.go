package main

import (
	"log"
)

func main() {

	myVar := 20

	switch {
	case myVar > 10:
		log.Println("Hi")
	case myVar > 11:
		log.Println("Hi")
	default:
		log.Println("default")
	}
}
