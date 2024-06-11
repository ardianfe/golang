package main

import (
	"fmt"
	"strings"
)

func main() {
	// Membuat string yang panjangnya 1300 byte
	longString := strings.Repeat("Hello, World!", 100)

	// Membuat buffer sebesar 1024 byte
	buffer := make([]byte, 1024)

	// Menyalin data dari string panjang ke buffer
	copied := copy(buffer, longString)

	// Menampilkan data yang ada di buffer
	fmt.Println("Buffer content:", string(buffer))
	fmt.Printf("Buffer length: %d bytes\n", len(buffer))
	fmt.Printf("Bytes copied: %d bytes\n", copied)
}
