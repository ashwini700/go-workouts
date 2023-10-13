// Create a map called fruits with fruit names as keys and their corresponding quantities as values.
// Add at least four different fruits and quantities to the map.
// Write code to iterate over the fruits map and print each fruit and its quantity.

package main

import "fmt"

type fruits struct {
	quantities float64
}

var m map[string]fruits

func main() {
	m = make(map[string]fruits)
	m["mango"] = fruits{10}
	m["apple"] = fruits{1}
	m["banana"] = fruits{100}
	m["guava"] = fruits{5}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

}
