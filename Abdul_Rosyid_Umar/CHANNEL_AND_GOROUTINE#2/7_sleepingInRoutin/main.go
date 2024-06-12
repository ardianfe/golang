package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine is sleeping")
			time.Sleep(1 * time.Second)
		}
	}()

	// Menunggu goroutine selesai
	time.Sleep(6 * time.Second)
}
