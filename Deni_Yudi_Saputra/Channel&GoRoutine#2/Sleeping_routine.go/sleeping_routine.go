package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Mulai")

	go func() {
		fmt.Println("Goroutine dimulai")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine bangun setelah 2 detik")
	}()

	// Menunggu goroutine selesai
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
