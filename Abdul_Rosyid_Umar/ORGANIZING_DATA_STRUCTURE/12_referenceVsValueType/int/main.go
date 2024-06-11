package main

import "fmt"

func main() {
	a := 10
	b := a // Salinan nilai a diberikan ke b
	b = 20 // Mengubah nilai b tidak mempengaruhi a

	fmt.Println(a) // Output: 10
	fmt.Println(b) // Output: 20
}
