package main

import (
	"fmt"
	"net/http"
)

// Handler interface mendefinisikan metode ServeHTTP
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
    handler := MyHandler{}
    http.Handle("/", handler)
    fmt.Println("Server is listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
