package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, Gopher Boy!")
	if err != nil {
		log.Println("error with writing the response: ", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, fmt.Sprintf("We are Egyptian Gophers, ready to GO! Today is : %s", getToday()))
	if err != nil {
		log.Println("error with writing the response: ", err)
	}
}

func getToday() string {
	today := time.Now().Weekday().String()
	return today
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting server on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
