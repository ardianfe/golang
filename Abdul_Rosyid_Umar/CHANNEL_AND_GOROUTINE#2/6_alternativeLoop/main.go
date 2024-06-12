package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat channel untuk mengakhiri loop goroutine
	done := make(chan bool)

	// Menjalankan goroutine berulang kali setiap detik menggunakan time.Tick
	go func() {
		ticker := time.Tick(1 * time.Second)
		for {
			select {
			case <-done:
				return
			case t := <-ticker:
				fmt.Println("Goroutine is running at", t)
			}
		}
	}()

	// Menunggu 5 detik sebelum mengakhiri goroutine
	time.Sleep(5 * time.Second)
	done <- true
}
