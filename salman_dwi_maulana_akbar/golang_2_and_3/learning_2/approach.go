package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height
}

func main() {
	shape := []Shape{
		Rectangle{Length: 10, Width: 5},
		Circle{Radius: 10},
		Triangle{Base: 10, Height: 5},
	}

	for _, s := range shape {
		fmt.Println("Area:", s.Area())
		fmt.Println("Perimeter:", s.Perimeter())
	}
}