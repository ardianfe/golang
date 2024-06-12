package main

import (
	"fmt"
	"net/http"
)

// Serial Link Checking adalah proses pemeriksaan status website secara berurutan.

func websiteChecker(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s might be down\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%s is up\n", url)
	} else {
		fmt.Printf("%s is returning %d\n", url, resp.StatusCode)
	}
}

func main() {
	website := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range website {
		websiteChecker(url)
	}
}
