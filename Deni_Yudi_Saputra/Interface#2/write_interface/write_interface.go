package main

import (
	"fmt"
	"os"
)

// Antarmuka Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// FileWriter adalah implementasi dari antarmuka Writer
type FileWriter struct {
	file *os.File
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	return fw.file.Write(p)
}

// BufferWriter adalah implementasi dari antarmuka Writer
// type BufferWriter struct {
//     buffer *bytes.Buffer
// }

// func (bw *BufferWriter) Write(p []byte) (n int, err error) {
//     return bw.buffer.Write(p)
// }

// Fungsi untuk menulis data menggunakan antarmuka Writer
func writeData(w Writer, data []byte) (int, error) {
	return w.Write(data)
}

func main() {
	// Data yang akan ditulis
	data := []byte("Hello, Kids\n")

	// Menulis ke file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fileWriter := &FileWriter{file: file}

	n, err := writeData(fileWriter, data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Successfully wrote %d bytes to file.\n", n)

	// Menulis ke buffer dalam memori
	// var buffer bytes.Buffer
	// bufferWriter := &BufferWriter{buffer: &buffer}

	// n, err = writeData(bufferWriter, data)
	// if err != nil {
	//     fmt.Println("Error writing to buffer:", err)
	//     return
	// }
	// fmt.Printf("Successfully wrote %d bytes to buffer.\n", n)
	// fmt.Println("Buffer contents:", buffer.String())
}
