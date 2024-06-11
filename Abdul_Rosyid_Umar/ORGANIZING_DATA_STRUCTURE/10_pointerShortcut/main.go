package main

import "fmt"

// Mendefinisikan struct bernama Book
type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	book := Book{"Go Programming", "John Doe", 29.99}
	p := &book // p adalah pointer ke book

	p.Price = 19.99                    // bisa langsung mengakses field melalui pointer
	fmt.Println("Updated book:", book) // Price berubah menjadi 19.99
}
