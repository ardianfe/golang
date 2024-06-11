package main

import (
	"fmt"
	"net/http"
)

// fungsi untuk website checker
func checkWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s might be down!\n", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%s is up and running!\n", url)
	} else {
		fmt.Printf("%s returned status code %d\n", url, resp.StatusCode)
	}
}

func main() {
	websites := []string{
		"http://localhost:8080/hello",
		"https://www.github.com",
		"https://www.stackoverflow.com",
	}

	for _, url := range websites {
		checkWebsite(url)
	}
}
