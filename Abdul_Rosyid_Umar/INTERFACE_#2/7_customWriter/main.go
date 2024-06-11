package main

import (
	"fmt"
)

// CustomWriter struct
type CustomWriter struct{}

// Implementasi metode Write
func (cw *CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Custom Writer:", string(p))
	return len(p), nil
}

func main() {
	// Membuat instance CustomWriter
	writer := &CustomWriter{}

	// Data yang akan ditulis menggunakan custom writer
	data := []byte("Hello, Custom Writer!")
	writer.Write(data)
}
