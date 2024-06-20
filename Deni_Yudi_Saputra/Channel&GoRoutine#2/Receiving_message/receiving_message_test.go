package receivingmessage

import (
	"fmt"
	"testing"
	"time"
)

func TestReceivingMessage(t *testing.T) {
	messages := make(chan string) // Membuat channel dengan tipe string

	// Goroutine untuk mengirim pesan ke channel
	go func() {
		time.Sleep(time.Second * 2)     // Menunggu 2 detik
		messages <- "Hello, Kiddsssss!" // Mengirim pesan ke channel
	}()

	// Menerima pesan dari channel/ receiving message
	msg := <-messages
	fmt.Println(msg)
}
