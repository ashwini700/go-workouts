package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4, 5}
	revSlice(number)
	fmt.Println(number)
}

func revSlice(num []int) {
	length := len(num)
	for i := 0; i < length/2; i++ {
		num[i], num[length-i-1] = num[length-i-1], num[i]
	}
}
