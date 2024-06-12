package main

import (
	"fmt"
)

func main() {
	// Deklarasi dan eksekusi function literal
	func() {
		fmt.Println("Hello from an anonymous function")
	}()

	// Menggunakan function literal dengan goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from a goroutine")
}
