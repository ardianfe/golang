package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
	Stock  int
}

func main() {
	// dengan struct
	book := Book{
		Title:  "Go Programming",
		Author: "John Doe",
		Price:  29.99,
		Stock:  100,
	}
	fmt.Println("Book:", book)

	// dengan map
	bookwithmap := map[string]interface{}{
		"Title":  "Go Programming",
		"Author": "John Doe",
		"Price":  29.99,
		"Stock":  100,
	}
	fmt.Println("Book:", bookwithmap)
}
