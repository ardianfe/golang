package main

import (
	"fmt"
	"strings"
)

func main() {
	// membuat string dengan variabel reader
	reader := strings.NewReader("Hello, World!")

	// membaca 5 bytes dari reader
	buffer := make([]byte, 5)
	_, err := reader.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the read data
	fmt.Println("Data Read:", string(buffer))
}
