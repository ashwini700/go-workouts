//Description: Write a program that doubles the values of all elements in a slice.
// Instructions:
// Create an integer slice called numbers with at least 6 different numbers.
// Write code to double the values of all elements in the numbers slice.
// Print the modified slice.

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	doubleValues(numbers)
	fmt.Println(numbers)
}

func doubleValues(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2 //doubleElements
	}
}
