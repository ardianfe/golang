package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Reader adalah interface untuk membaca data
type Reader interface {
	Read(p []byte) (n int, err error)
}

// FileReader adalah implementasi Reader untuk membaca dari file
type FileReader struct {
	file *os.File
}

// NewFileReader mengembalikan instance baru FileReader untuk file tertentu
func NewFileReader(filename string) (*FileReader, error) {
	// path buat baca file dari folder lain 
	absPath, err := filepath.Abs(filename)
	if err != nil{
		return nil, err
	}
	
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	return &FileReader{file: file}, nil
}

// Read membaca data dari file menggunakan FileReader
func (fr *FileReader) Read(p []byte) (n int, err error) {
	return fr.file.Read(p)
}

// Close menutup file setelah selesai digunakan
func (fr *FileReader) Close() {
	fr.file.Close()
}


func main() {
	// menggunakan file path nya 
	filepath := "../more_reader_interface/contoh.txt"
	// Menggunakan FileReader
	fileReader, err := NewFileReader(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	buf := make([]byte, 1024)
	n, err := fileReader.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("FileReader: Read %d bytes\n", n)
	fmt.Printf("FileReader: Data read: %s\n", buf[:n])

	
}
