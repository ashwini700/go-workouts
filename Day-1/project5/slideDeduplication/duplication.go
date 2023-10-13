package main

import "fmt"

func removeDups(nums []int) []int {
	//to store distinct elements
	distinct := map[int]bool{}
	result := []int{}

	//adding looped elements which are unique
	for v := range nums {
		if distinct[nums[v]] == false {
			distinct[nums[v]] = true
			result = append(result, nums[v])
		}
	}

	return result
}

func main() {
	var numbers = []int{1, 2, 2, 3, 4, 9, 0, 2, 5, 6, 1, 7, 8, 9, 3}
	duplicate := removeDups(numbers)
	fmt.Println("Original:", numbers)
	fmt.Println("Deduplicate:", duplicate)
}
