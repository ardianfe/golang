package main

import (
	"fmt"
)

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter interface
type ReadWriter interface {
	Reader
	Writer
}

// Struct yang mengimplementasikan keduanya Reader dan Writer
type MyReadWriter struct {
	data []byte
}

// Implementasi Read method
func (rw *MyReadWriter) Read(p []byte) (n int, err error) {
	copy(p, rw.data)
	return len(rw.data), nil
}

// Implementasi Write method
func (rw *MyReadWriter) Write(p []byte) (n int, err error) {
	rw.data = append(rw.data, p...)
	return len(p), nil
}

func main() {
	rw := &MyReadWriter{}

	// tulis data
	rw.Write([]byte("Hello, World!"))

	// baca data
	buffer := make([]byte, 12)
	rw.Read(buffer)
	fmt.Println("Data:", string(buffer))
}
