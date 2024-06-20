package main

import "fmt"

// interface 1
type Reader interface {
    Read() string
}

// interface 3
type Writer interface {
    Write(text string)
}

// gabungan interface
type ReadWriter interface {
    Reader
    Writer
}

// =================

// Tipe yang mengimplementasikan ReadWriter
type MyReadWriter struct {
    content string
}

func (rw *MyReadWriter) Read() string {
    return rw.content
}

func (rw *MyReadWriter) Write(text string) {
    rw.content = text
}

func main() {
    // Membuat instance dari MyReadWriter
    var rw ReadWriter = &MyReadWriter{}

    // Menggunakan metode Write dari antarmuka Writer
    rw.Write("Hello, Go!")
    
    // Menggunakan metode Read dari antarmuka Reader
    fmt.Println(rw.Read()) // Output: Hello, Go!
}
