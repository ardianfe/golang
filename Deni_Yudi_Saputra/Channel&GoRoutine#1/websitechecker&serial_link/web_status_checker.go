package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkWebsite(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Mengurangi hitungan WaitGroup sebesar 1 ketika fungsi selesai
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s tidak dapat diakses!\n", url)
		return
	}
	fmt.Printf("Sukses: %s dapat diakses, kode status: %d\n", url, resp.StatusCode)
	resp.Body.Close()
}

func main() {
	urls := []string{
		"https://www.geeksforgeeks.org",
		"https://docs.google.com/document/d/1ICCywZu3FbgqsisDnjFDOYJkUvt0Xno98evaR3pELcI/edit#heading=h.c9kwfrbobt0q",
		"https://www.thisurldoesnotexist12345.com", // Tidak dapat diakses
		"http://www.invalid-url.com",               // Tidak dapat diakses
		"https://www.fakeurl.notreal",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1) // Menambahkan satu goroutine ke WaitGroup
		go checkWebsite(url, &wg)
	}

	wg.Wait()
}
