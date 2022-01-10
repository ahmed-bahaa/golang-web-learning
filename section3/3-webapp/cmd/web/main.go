package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ahmed-bahaa/thirdwebapp/pkg/handlers"
)

const portNumber = ":8080"

func getToday() string {
	today := time.Now().Weekday().String()
	return today
}

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting server on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
