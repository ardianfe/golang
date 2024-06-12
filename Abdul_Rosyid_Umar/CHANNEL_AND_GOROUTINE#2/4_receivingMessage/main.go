package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	statusChan := make(chan string)

	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range urls {
		go checkWebsite(url, statusChan)
	}

	// Menggunakan select statement untuk menerima pesan dari channel
	for i := 0; i < len(urls); i++ {
		select {
		case msg := <-statusChan: // menerima channel
			fmt.Println(msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout waiting for response")
		}
	}
}

func checkWebsite(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s is down!", url) // menerima message
		return
	}
	time.Sleep(2 * time.Second)                                         // Simulasi waktu pemrosesan
	c <- fmt.Sprintf("%s is up! Status Code: %d", url, resp.StatusCode) //menerima message
}
