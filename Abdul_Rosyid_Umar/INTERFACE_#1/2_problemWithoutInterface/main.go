package main

// Tanpa interface, kita harus menulis fungsi terpisah untuk
// Setiap tipe data, yang membuat kode menjadi redundan dan sulit untuk dikelola

import "fmt"

// Struct representing a Circle
type Circle struct {
	Radius float64
}

// Method to calculate the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Struct representing a Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to print the area of a circle
func PrintCircleArea(c Circle) {
	fmt.Println("Circle Area:", c.Area())
}

// Function to print the area of a rectangle
func PrintRectangleArea(r Rectangle) {
	fmt.Println("Rectangle Area:", r.Area())
}

func main() {
	// Create a circle
	circle := Circle{Radius: 5}
	// Create a rectangle
	rectangle := Rectangle{Width: 3, Height: 4}

	// Print the area of the circle
	PrintCircleArea(circle)

	// Print the area of the rectangle
	PrintRectangleArea(rectangle)
}
