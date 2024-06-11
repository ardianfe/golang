package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Struct untuk menyimpan data JSON
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler function untuk endpoint /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Handler function untuk /submit endpoint
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Received: Name = %s, Age = %d", data.Name, data.Age)
		w.Write([]byte(response))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
