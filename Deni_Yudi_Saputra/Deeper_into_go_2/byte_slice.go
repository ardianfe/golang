package main

import "fmt"

func main() {
	data := []byte{'H', 'e', 'l', 'l', 'o', '!'}

	// slice := make([]byte, 5)      // Membuat slice dengan panjang 5 dan kapasitas 5 (awalnya diisi dengan nol)
	// slice1 := make([]byte, 5, 10) // Membuat slice dengan panjang 5 dan kapasitas 10

	fmt.Println(string(data[2:5]))
	fmt.Println(data)
	
}
