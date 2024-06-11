package main

import "fmt"

// mendefinisikan interface
type Shape interface {
	Area() float64
}

// Struct yang merepresentasikan lingkaran
type Circle struct {
	Radius float64
}

// Method untuk menghitung area atau luas dari lingkaran
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// membuat lingkaran
	circle := Circle{Radius: 5}

	// Assign lingkaran ke Shape interface
	var shape Shape
	shape = circle

	// memanggil method via Shape interface
	fmt.Println("Circle Area:", shape.Area())
}
