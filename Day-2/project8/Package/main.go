package main

import (
	"Package/model"
	"Package/person"
)

func main() {

	p := model.Person{
		Name: "Ashwini",
		Age:  23,
	}
	person.PrintPerson(p)
}
