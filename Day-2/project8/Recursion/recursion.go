package main

import "fmt"

func fact(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return i * fact(i-1)
}
func main() {
	fmt.Println(fact(3))
}
