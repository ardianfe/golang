package main

import (
	"fmt"
	"testing"
	"time"
)

// dalam go berarti ketika blocking channel berarti bahwa pengiriman atau penerimaan akan berhenti atau mengunggu sampai operasi pada goroutine lain selesai

func TestBlocking(t *testing.T) {
	channel := make(chan int)

	// goroutine untuk mengirim data ke channel

	go func() {
		fmt.Println("Mengirim nilai 42..")
		channel <- 42 //kirim nilai ke channel
		fmt.Println("Nilai 42 terkirim")
	}()

	// sampai sini diblokir sampai ada penerimanya

	time.Sleep(5 * time.Second) // Simulasi penundaan
	fmt.Println("Menerima nilai dari channel...")
	value := <-channel // Menerima data dari channel
	fmt.Println("Diterima:", value)
}
