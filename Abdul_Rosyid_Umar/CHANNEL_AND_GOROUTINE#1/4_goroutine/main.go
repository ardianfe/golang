package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range urls {
		// Menjalankan checkWebsite sebagai goroutine
		go checkWebsite(url)
	}

	// Tunggu selama 5 detik agar semua goroutine selesai
	// program mau tidak mau dihentikan setelah 5 detik meskipun go routine belum selesai
	// solusinya adalah menggunakan channel agar tidak memakai time.sleep ini
	time.Sleep(5 * time.Second)
}

func checkWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is down!\n", url)
		return
	}
	fmt.Printf("%s is up! Status Code: %d\n", url, resp.StatusCode)
}
