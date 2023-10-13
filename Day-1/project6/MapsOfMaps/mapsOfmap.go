// n: Create a program that uses a map of maps to store and retrieve data.
// Instructions:
// Create a map called studentData where each student's name is the key, and the value is another map containing information such as Age, Grade, and City.
// Add at least three students to the studentData map with their respective information.
// Retrieve and print the details of each student.

package main

import "fmt"

func main() {
	studentData := make(map[string]map[string]interface{})
	studentData["Ashwini"] = map[string]interface{}{
		"age":   12,
		"grade": "A",
		"city":  "Bijapur",
	}
	studentData["Shaine Nigam"] = map[string]interface{}{
		"age":   23,
		"grade": "B",
		"city":  "kerala",
	}
	studentData["Salman"] = map[string]interface{}{
		"age":   50,
		"grade": "A",
		"city":  "Mumbai",
	}

	fmt.Println("studentData of Ashwini:", studentData["Ashwini"])
	fmt.Println("studentData of Shaine:", studentData["Shaine Nigam"])
	fmt.Println("studentData of Salman:", studentData["Salman"])

}
