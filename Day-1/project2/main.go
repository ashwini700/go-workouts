// main.go

package main

import (
	"fmt"
	"project1/calculator"
)

func main() {
	n := 5
	square := calculator.Square(n)
	fmt.Printf("Square of %d: %d\n", n, square)

	sum := calculator.Sum(10, 5)
	fmt.Printf("Sum: %d\n", sum)

	diff := calculator.Difference(10, 5)
	fmt.Printf("Difference: %d\n", diff)

	product := calculator.Product(8, 3)
	fmt.Printf("Product: %d\n", product)

	q, r := calculator.QuotAndRem(15, 4)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}
