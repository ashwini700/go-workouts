package arrays

import "fmt"

func Arr(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
