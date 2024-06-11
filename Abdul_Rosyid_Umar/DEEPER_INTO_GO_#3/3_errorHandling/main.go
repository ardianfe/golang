package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("nonexistent.txt")
	if err != nil { // eror handling
		fmt.Println("Error:", err) // eror akan muncul karena nonexistent.txt tidak ada di direktori
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
