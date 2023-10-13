// Description: Write a program that finds and prints the index of a specific number in a slice.
// Instructions:
// Create an integer slice called numbers with at least 6 different numbers.
// Choose a number to search for in the numbers slice.
// Write code to find the index of the chosen number. Return -1 if the element is not found.
// Print the index (or a message if the number is not found).

package main

import "fmt"

func main() {
	numbers := []int{12, 15, 16, 17, 18, 10}
	searchNum := 40
	// fmt.Printf("search number : %d ", numbers[7])
	index := -1
	for i, num := range numbers {
		if num == searchNum {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Printf("index of %d is: %d\n", searchNum, index)
	} else {
		fmt.Printf("%d not found in slice.\n", searchNum)
	}
}
