// Description: Create a simple program that uses a map to store and retrieve information.
// Instructions:
// Create a map called studentGrades with student names as keys and their corresponding grades as values.
// Add at least three students to the map with their grades.
// Print the grades of each student.
// Delete one student's entry from the map.
// Now print the students map again.
// Note: student grades are float values between 0 to 10. Example 5.6, 9.2 etc

package main

import "fmt"

type studentGrades struct {
	// studentNames string
	grades float64
}

var m map[string]studentGrades

func main() {
	m = make(map[string]studentGrades)
	m["Ashwini"] = studentGrades{9.5}
	m["Ganii"] = studentGrades{9.0}
	m["Ajay"] = studentGrades{7.0}
	fmt.Println(m["Ashwini"])
	fmt.Println(m["Ganii"])
	fmt.Println(m["Ajay"])

	key := "Ganii"
	delete(m, key)

	fmt.Println(key, m)
}
