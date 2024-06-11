package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Membuat file baru
	err := ioutil.WriteFile("source.txt", []byte("Hello golangasdfasd"), 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Membuka file sumber
	srcFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Error membuka file sumber:", err)
		return
	}
	defer srcFile.Close()

	// Membuat file tujuan
	dstFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error membuat file tujuan:", err)
		return
	}
	defer dstFile.Close()

	// Menyalin data dari file sumber ke file tujuan
	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Error menyalin data:", err)
		return
	}
	fmt.Printf("%d byte telah disalin dari source.txt ke destination.txt\n", bytesCopied)
}
