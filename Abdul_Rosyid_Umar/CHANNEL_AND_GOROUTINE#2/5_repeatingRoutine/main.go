package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat channel untuk mengakhiri loop goroutine
	done := make(chan bool)

	// Menjalankan goroutine berulang kali setiap detik
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Goroutine is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Menunggu 5 detik sebelum mengakhiri goroutine
	time.Sleep(5 * time.Second)
	done <- true
}
