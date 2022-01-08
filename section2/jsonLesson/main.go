package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
	{
		"first": "Clark",
		"last": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first": "Bruce",
		"last": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
	]`

	var unmarshaled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshaled)
	if err != nil {
		log.Println("Error with unmrasal", err)
	}

	log.Printf("Unmarshaled value %v", unmarshaled)

	var mySlice []Person
	p1 := Person{
		FirstName: "AHmed",
		LastName:  "Baha",
		HairColor: "Gray",
		HasDog:    true,
	}

	p2 := Person{
		FirstName: "SILO",
		LastName:  "IW",
		HairColor: "Blue",
		HasDog:    false,
	}
	mySlice = append(mySlice, p1)
	mySlice = append(mySlice, p2)

	res, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		log.Println("Error with Marhaling hte json: ", err)
	}

	fmt.Println("Result:", string(res))

}
