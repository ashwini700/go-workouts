package main

import (
	"fmt"
	calculates "sumAvg/calculate"
)

func main() {

	// //slices
	// fruits := []string{"apple", "banana", "jackfruit", "pineapple", "mango"}
	// slices.fruits(fruits)

	//append , add ,removal , len , cap

	// var number []int
	// number = append.Add(number, 5)
	// number = append.Add(number, 10)
	// number = append.Add(number, 15)
	// number = append.Add(number, 20)
	// number = append.Add(number, 25)

	// // length and capacity
	// fmt.Printf("Length: %d, Capacity: %d\n", len(number), cap(number))

	// // Remove number 15
	// number = append.Remove(number, 15)
	// //length and capacity after removal
	// fmt.Printf("len: %d, Cap: %d\n", len(number), cap(number))

	//calculation of sum and average
	num := []int{10, 20, 30, 40, 50}

	// sum of numbers
	sum := calculates.Sum(num)
	fmt.Printf("Sum: %d\n", sum)

	// Calculate and print the average of the numbers
	average := calculates.Average(num)
	fmt.Printf("Average: %.2f\n", average)

	calculates.Average(num)
}
