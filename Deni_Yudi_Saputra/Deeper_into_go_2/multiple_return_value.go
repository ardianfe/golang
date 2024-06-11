package main

import (
	"fmt"
)

// func hitung(d float64) (float64, float64) {
// 	area := math.Pi * math.Pow(d*2, 2) //luas
// 	kel := math.Pi * d                 //keliling

// 	return area, kel

// }

func hitung2(a, b int) (int, int) {
	bagi := a / b
	kurang := a - b

	return bagi, kurang
}



func main() {
	// var diameter float64 = 15

	// area, kel := hitung(diameter)

	// fmt.Printf("Luas=", area)
	// fmt.Printf("Keliling=", kel)

	bagi, kurang := hitung2(10, 2)
	fmt.Printf("Bagi=", bagi)
	fmt.Printf("Kurang=", kurang)

}
