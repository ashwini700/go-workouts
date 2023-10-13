package main

import (
	"fmt"
	maxelement "maximum/max_element"
)

func main() {
	num := []int{10, 18, 17, 20, 5}
	max := maxelement.Max(num)
	fmt.Printf("The max element in slice : %d\n", max)
}
