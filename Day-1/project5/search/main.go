package main

import (
	"fmt"
	"os"
	"strconv"
)

func search(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func main() {

	numbers := []int{40, 50, 60, 30, 80}
	if len(os.Args) < 2 {
		fmt.Println(" number to search.")
		return
	}

	searchStr := os.Args[1]
	searchNum, err := strconv.Atoi(searchStr)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	found := search(numbers, searchNum)
	if found {
		fmt.Printf("Number %d was found in the slice.\n", searchNum)
	} else {
		fmt.Printf("Number %d was not found in the slice.\n", searchNum)
	}
}
