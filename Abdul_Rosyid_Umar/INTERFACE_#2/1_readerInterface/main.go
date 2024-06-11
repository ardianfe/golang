package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, World!") // Membuat reader dari string
	buffer := make([]byte, 8)                    // Membuat buffer sebesar 8 byte

	for {
		n, err := reader.Read(buffer) // Membaca data ke dalam buffer
		if err != nil {
			break // Keluar dari loop jika terjadi error atau EOF
		}
		fmt.Println("Read:", string(buffer[:n])) // Menampilkan data yang dibaca
	}
}
