package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	//sum of even num

	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}

	fmt.Printf("Sum of even nums %d\n", sum)
}
