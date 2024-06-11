package main

import (
	"fmt"
	"net/http"
)

// Fungsi untuk memeriksa status website
func checkWebsite(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s might be down!", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		ch <- fmt.Sprintf("%s is up and running!", url)
	} else {
		ch <- fmt.Sprintf("%s returned status code %d", url, resp.StatusCode)
	}
}

func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
	}

	// Membuat channel
	ch := make(chan string)

	// Menjalankan goroutine untuk setiap website
	for _, url := range websites {
		go checkWebsite(url, ch)
	}

	// Menerima hasil dari setiap goroutine
	for i := 0; i < len(websites); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Main function finished")
}
