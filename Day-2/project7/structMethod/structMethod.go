// Description: Create a struct with methods for common operations.
// Instructions:
// Define a struct called Rectangle with fields for Width and Height.
// Create methods for calculating the area and perimeter of a rectangle on this struct.
// Create an instance of the Rectangle struct and use the methods to calculate its area and perimeter.
// Print the results.

package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	Rectangle := Rectangle{
		Width:  10,
		Height: 10,
	}
	fmt.Println(Rectangle.Area())
	fmt.Println(Rectangle.perimeter())
}
