package calculates

func Sum(num []int) int {
	sum := 0
	for _, num := range num {
		sum += num
	}
	return sum
}

func Average(num []int) float64 {
	if len(num) == 0 {
		return 0
	}

	sum := Sum(num)
	return float64(sum) / float64(len(num))
}
