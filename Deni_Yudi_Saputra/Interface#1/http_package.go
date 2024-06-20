package main

import (
	"fmt"
	"net/http"
)

// http.Handler merupakan interface utamanya untuk menangani berbagai operasi HTTP

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// ServeHTTP: Metode ini dipanggil untuk menangani setiap permintaan HTTP. Parameter ResponseWriter digunakan untuk mengirim tanggapan kembali ke klien, dan *Request berisi semua informasi tentang permintaan HTTP yang diterima.

// Implementasi

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a custom handler!")
}

// handle func
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a handler function!")
}

func main() {
	handler := MyHandler{}
	http.Handle("/", handler)

	//handlefunc
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server berlajan pada port 8080")
	http.ListenAndServe(":8080", nil)

}
