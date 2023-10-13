package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	countEven, countOdd := countEvenOdd(numbers)
	fmt.Printf("Even numbers: %d\n", countEven)
	fmt.Printf("Odd numbers: %d\n", countOdd)
}

func countEvenOdd(numbers []int) (countEven, countOdd int) {
	for _, num := range numbers {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	return countEven, countOdd
}
