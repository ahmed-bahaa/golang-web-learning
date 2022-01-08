package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, Gopher Boy!")
		if err != nil {
			log.Println("error with writing the response: ", err)
		}
		log.Println("Number of bytes written: ", n)
	})
	_ = http.ListenAndServe(":8080", nil)

}
