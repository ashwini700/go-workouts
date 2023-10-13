package main

import (
	"fmt"
	"project2/temperature"
)

func main() {
	var fahrenheit float64
	fmt.Print("Enter Fahrenheit temperature: ")
	_, err := fmt.Scanf("%f", &fahrenheit)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	celsius := temperature.FahrenToCel(fahrenheit)

	fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
}
