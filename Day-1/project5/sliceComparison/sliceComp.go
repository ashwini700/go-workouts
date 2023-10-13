package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 7}

	isEqual := true
	if len(slice1) != len(slice2) {
		isEqual = false
	} else {
		for i := 0; i < len(slice1); i++ {
			if slice1[i] != slice2[i] {
				isEqual = false
				break
			}
		}
	}

	if isEqual {
		fmt.Println("Both slices are equal")
	} else {
		fmt.Println("slices are not equal.")
	}
}
