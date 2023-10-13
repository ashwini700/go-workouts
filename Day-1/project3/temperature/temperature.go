package temperature

func FahrenToCel(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
