package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func readFromReader(reader io.Reader) {
	// Buffer untuk menampung data yang dibaca
	buffer := make([]byte, 10)

	for {
		// Membaca data ke dalam buffer
		n, err := reader.Read(buffer)
		if err == io.EOF {
			// Akhir dari sumber data, keluar dari loop
			break
		}
		if err != nil {
			// Menangani error lainnya
			fmt.Println("Error membaca:", err)
			break
		}
		fmt.Println("Baca:", string(buffer[:n]))
	}
}

func main() {
	// Membuat file baru
	err := ioutil.WriteFile("example.txt", []byte("Hello golangasdfasd"), 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Membaca dari file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()
	fmt.Println("Membaca dari file:")
	readFromReader(file)

	// Membaca dari string
	strReader := strings.NewReader("Hello, Go!")
	fmt.Println("\nMembaca dari string:")
	readFromReader(strReader)
}
