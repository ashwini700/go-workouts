package append

import "fmt"

// Add and append numbers
func Add(number []int, num int) []int {
	number = append(number, num)
	fmt.Printf("after addition  %d: %v\n", num, number)
	return number
}

func Remove(number []int, num int) []int {
	for i, n := range number {
		if n == num {
			number = append(number[:i], number[i+1:]...)
			break
		}
	}
	fmt.Printf("after removal of numbers %d: %v\n", num, number)
	return number
}
