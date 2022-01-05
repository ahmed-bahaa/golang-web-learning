package main

import (
	"log"

	"github.com/ahmed-bahaa/my-random-intchan/helpers"
)

const rangeNum = 1000

func getRandomNumber(chanInt chan int) {
	i := 0
	for i < 3 {
		myValue := helpers.GetRandomInt(rangeNum)
		chanInt <- myValue
		i++
	}

}

func main() {

	chanInt := make(chan int)
	defer close(chanInt)

	go getRandomNumber(chanInt)

	for i := 0; i < 3; i++ {
		theRandomNumber := <-chanInt
		log.Println(theRandomNumber)
	}

}
