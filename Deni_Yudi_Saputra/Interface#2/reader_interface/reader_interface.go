package main

import (
	"fmt"
)

// Reader adalah interface untuk membaca data.
type Reader interface {
	Read(p []byte) (n int, err error)
}

// StringBuffer adalah tipe data yang mengimplementasikan Reader interface
type StringBuffer struct {
	data string
}

// Implementasi metode Read untuk StringBuffer
func (sb *StringBuffer) Read(p []byte) (n int, err error) {
	// Copy string data into byte slice p
	n = copy(p, []byte(sb.data))
	if n == 0 {
		return 0, fmt.Errorf("end of data")
	}
	return n, nil
}

func main() {
	// Membuat objek StringBuffer
	sb := &StringBuffer{data: "Slebeww ya"}

	// Membuat byte slice untuk menampung hasil pembacaan
	data := make([]byte, len(sb.data))

	// Menggunakan objek StringBuffer untuk membaca data
	n, err := sb.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Mencetak data yang telah dibaca
	fmt.Printf("Read %d bytes: %s\n", n, data)
}
