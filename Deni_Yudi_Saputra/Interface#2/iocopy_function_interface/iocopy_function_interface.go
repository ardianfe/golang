package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// io.Copy sendiri ialah fungsinya untuk menyalin data dari sumber (io.Reader) ke tujuan (io.Writer) dalam konteks transfer data

// Menyalin file

func menyalinFile() {
	srcFile, err := os.Open("../more_reader_interface/contoh.txt")
	if err != nil {
		fmt.Println("Error membuka file sumber:", err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error membuat file tujuan:", err)
		return
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Error menyalin data:", err)
		return
	}

	fmt.Printf("Berhasil menyalin %d byte dari source.txt ke destination.txt\n", bytesCopied)
}

// Menyalin HTTP Respone ke File

func menyalinHttpFile(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error melakukan permintaan HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Periksa status kode HTTP
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("gagal mengunduh file: %s", resp.Status)
	}

	// Periksa header tipe konten
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/zip" {
		return fmt.Errorf("konten yang diterima bukan file ZIP, melainkan: %s", contentType)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error membuat file: %v", err)
	}
	defer outFile.Close()

	bytesCopied, err := io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error menyalin data: %v", err)
	}

	fmt.Printf("Berhasil menyalin %d byte dari response HTTP ke %s\n", bytesCopied, filename)
	return nil
}

func main() {
	menyalinFile()
	
	url := "https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-zip-file.zip" // Gantilah dengan URL file ZIP yang benar
	filename := "file.zip"                        // Nama file yang sesuai

	if err := menyalinHttpFile(url, filename); err != nil {
		fmt.Println("Error:", err)
	}
}
