package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Membuat file baru
	err := ioutil.WriteFile("source.txt", []byte("Hello copyio"), 0644)
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

	// Menyalin data dari file sumber ke file tujuan menggunakan fungsi customCopy
	bytesCopied, err := customCopy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Error menyalin data:", err)
		return
	}
	fmt.Printf("%d byte telah disalin dari source.txt ke destination.txt\n", bytesCopied)
}

// customCopy menyalin data dari src (Reader) ke dst (Writer)
func customCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	// Membuat buffer untuk menampung data yang dibaca
	buffer := make([]byte, 32*1024) // 32KB buffer

	for {
		// Membaca data dari src ke buffer
		n, readErr := src.Read(buffer)
		if n > 0 {
			// Menulis data dari buffer ke dst
			nw, writeErr := dst.Write(buffer[:n])
			if nw > 0 {
				written += int64(nw)
			}
			if writeErr != nil {
				err = writeErr
				break
			}
			if n != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if readErr != nil {
			if readErr != io.EOF {
				err = readErr
			}
			break
		}
	}

	return written, err
}
