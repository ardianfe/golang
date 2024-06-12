// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan int) // Membuat unbuffered channel

// 	go func() { // Memulai goroutine baru
// 		time.Sleep(2 * time.Second) // Tidur selama 2 detik
// 		ch <- 42                    // Mengirim nilai 42 ke channel
// 	}()

// 	select {
// 	case val := <-ch: // Mencoba menerima nilai dari channel
// 		fmt.Println("Received:", val)
// 	case <-time.After(1 * time.Second): // Menunggu 1 detik untuk timeout
// 		fmt.Println("Timeout!")
// 	} // output akan timeout atau channel terblokir karena waktu yang dibutuhkan goroutine melebihi 1 detik di salah satu case select
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
	time.Sleep(9 * time.Second) // Simulasi waktu pemrosesan
	c <- fmt.Sprintf("%s is up! Status Code: %d", url, resp.StatusCode)
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
		go checkWebsite(url, statusChan)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
	//Channel bersifat blocking, jadi goroutine main akan menunggu
	//sampai checkWebsite mengirim data ke channel sebelum melanjutkan.
}
