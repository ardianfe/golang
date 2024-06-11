package main

import "fmt"

// Mendefinisikan struct bernama Book
type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	var p *Book // p adalah null pointer (belum diinisialisasi)

	// fmt.Println(p.Price) // ini akan menyebabkan panic karena p adalah null pointer

	book := Book{"Go Programming", "John Doe", 29.99}
	p = &book            // inisialisasi pointer p
	fmt.Println(p.Price) // aman karena p telah diinisialisasi
}
