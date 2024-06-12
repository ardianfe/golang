package main

import (
	"fmt"
	"net/http"
)

// Printing Site Status adalah bagian dari program yang mencetak status dari website yang diperiksa.

func websiteStatusCheck(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s might be down\n", url)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%s is up!\n", url)
	} else {
		fmt.Printf("%s returned status: %d\n", url, resp.StatusCode)
	}
}

func main() {
	website := []string{
		"http://localhost:3000",
		"http://github.com",
		"http://stackoverflow.com",
	}

	for _, url := range website {
		websiteStatusCheck(url)
	}
}
