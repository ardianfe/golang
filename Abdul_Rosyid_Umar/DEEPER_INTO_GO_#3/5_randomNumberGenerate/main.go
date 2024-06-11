package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) // Generates random number dari 0 sampe 99
	fmt.Println("Random number:", randomNumber)
}
