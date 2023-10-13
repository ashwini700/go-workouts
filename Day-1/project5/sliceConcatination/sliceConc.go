// Slice Concatenation

// Description: Create a program that concatenates two slices into a single slice.
// Instructions:
// Create two integer slices, slice1 and slice2, each with at least 3 different numbers.
// Write code to concatenate slice2 to the end of slice1.
// Print the concatenated slice.

package main

import "fmt"

func main() {
	sl1 := []int{8, 6, 5} //slice1
	sl2 := []int{4, 3, 7} //slice2

	sliceConc := concat(sl1, sl2)
	fmt.Println("Concatenated Slices:", sliceConc)
}

func concat(slice1, slice2 []int) []int {
	concatedSlice := append(slice1, slice2...) //append to conc slices
	return concatedSlice
}
