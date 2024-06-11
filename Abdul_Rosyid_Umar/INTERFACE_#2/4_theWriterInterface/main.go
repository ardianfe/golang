package main

import (
	"fmt"
	"os"
)

// Mendefinisikan interface Writer yang memiliki metode Write
type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	// Membuat file untuk menulis
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error membuat file:", err)
		return
	}
	defer file.Close()

	// Data yang akan ditulis
	data := []byte("Hello, Go!")

	// Menulis data ke file
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Error menulis ke file:", err)
		return
	}
	fmt.Printf("%d byte telah ditulis ke file\n", n)
}
