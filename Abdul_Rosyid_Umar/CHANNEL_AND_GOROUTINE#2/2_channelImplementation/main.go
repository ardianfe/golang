// package main

// import "fmt"

// func main() {
// 	ch := make(chan int)

// 	go func() { // menggunakan goroutine
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 		close(ch) // Menutup channel setelah selesai mengirim data
// 	}()

// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}
// }

package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebsite(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s is down!", url)
		return
	}
	c <- fmt.Sprintf("%s is up! Status Code: %d", url, resp.StatusCode)
	time.Sleep(1 * time.Second) // Simulasi waktu pemrosesan
}

func main() {
	// Membuat channel untuk komunikasi antar goroutine
	statusChan := make(chan string)

	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range urls {
		// Menjalankan checkWebsite sebagai goroutine
		go checkWebsite(url, statusChan)
	}

	// Menerima status dari channel untuk setiap URL
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
