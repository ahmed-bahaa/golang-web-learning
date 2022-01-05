package main

import (
	"log"

	"github.com/ahmed-bahaa/my-interfaces/helpers"
)

type animal interface {
	Say() string
	GetNoLeg() int
}

type cat struct {
	Name string
	Legs int
}

type dog struct {
	Name string
	Legs int
}

func main() {
	c := cat{
		Name: "zeus",
		Legs: 4,
	}

	DefineAnimal(&c)

	d := dog{
		Name: "mars",
		Legs: 5,
	}

	DefineAnimal(&d)
	helpers.SaySomething()

}

func DefineAnimal(a animal) {
	log.Println(a.GetNoLeg(), a.Say())
}

func (c *cat) Say() string {
	return "Meow"
}
func (c *cat) GetNoLeg() int {
	return c.Legs
}

func (d *dog) Say() string {
	return "Park"
}
func (d *dog) GetNoLeg() int {
	return d.Legs
}
