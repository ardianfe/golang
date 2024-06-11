package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := p1  // Salinan dari p1 diberikan ke p2
	p2.X = 10 // Mengubah p2 tidak mempengaruhi p1

	fmt.Println(p1) // Output: {1, 2}
	fmt.Println(p2) // Output: {10, 2}
}
