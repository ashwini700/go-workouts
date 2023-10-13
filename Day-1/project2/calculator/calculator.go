package calculator

// Square
func Square(n int) int {
	return n * n
}

// Sum
func Sum(a, b int) int {
	return a + b
}

// Difference
func Difference(a, b int) int {
	return a - b
}

// Product
func Product(a, b int) int {
	return a * b
}

// Quotient And Remainder
func QuotAndRem(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
