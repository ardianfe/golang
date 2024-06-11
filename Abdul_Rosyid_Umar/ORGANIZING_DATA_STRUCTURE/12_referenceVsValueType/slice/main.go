package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a    // b merujuk ke slice yang sama dengan a
	b[0] = 10 // Mengubah b juga mempengaruhi a

	fmt.Println(a) // Output: [10, 2, 3]
	fmt.Println(b) // Output: [10, 2, 3]
}
