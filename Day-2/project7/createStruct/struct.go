// Description: Define a struct for representing an Employee and create instances of it.
// Instructions:
// Define a struct called Employee with fields for Id, Name, Age, and City.
// Create two instances of the Employee struct with different values.
// Print the details of each employee.

package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	Employee1 := Employee{
		Id:   1,
		Name: "Ashwini",
		Age:  23,
		City: "Bijapur",
	}
	Employee2 := Employee{
		Id:   2,
		Name: "Prabhas",
		Age:  43,
		City: "AP",
	}
	fmt.Println("Employee1:", Employee1)
	fmt.Println("Employee2:", Employee2)

}
