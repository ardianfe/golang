package main

import "fmt"

// mendefinisikan interface shape dengan metode area
type Shape interface {
	Area() float64
}

// Struct untuk merepresentasikan circle
type Circle struct {
	JariJari float64 // atau disebut juga radius
}

// Metode untuk menghitung area khusus untuk circle
func (c Circle) Area() float64 {
	return 3.14 * c.JariJari * c.JariJari
}

// Struct yang merepresentasikan persegi panjang
type Rectangle struct {
	Width  float64
	Height float64
}

// Metode untuk menghitung area khusus untuk persegi panjang
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %v\n", s.Area())
}

func main() {
	// buat lingkaran
	circle := Circle{JariJari: 5}
	// buat persegi panjang
	rectangle := Rectangle{Width: 3, Height: 4}

	// print area lingkaran
	printArea(circle)

	// print area persegi panjang
	printArea(rectangle)
}
