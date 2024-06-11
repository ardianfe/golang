package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// membuat file baru
	err := ioutil.WriteFile("example.txt", []byte("Hello golangggg"), 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// buffer untuk menyimpan file yang dibaca
	buffer := make([]byte, 10)

	for {
		// membaca data dan masukan ke dalam buffer
		n, err := file.Read(buffer)
		if err == io.EOF {
			// keluar dari loop jika mencapai akhir data (EOF)
			break
		}
		if err != nil {
			// Handle eror lain
			fmt.Println("Error reading file:", err)
			break
		}
		// fmt.Println("Read:", string(buffer[:n]))
		fmt.Printf("%s", buffer[:n])
	}
}
