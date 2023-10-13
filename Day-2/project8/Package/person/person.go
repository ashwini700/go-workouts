package person

import (
	"Package/model"
	"fmt"
)

func PrintPerson(p model.Person) {

	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
