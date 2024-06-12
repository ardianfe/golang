package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to produce different results each run
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random integer between 0 and 99
	randomNumber := rng.Intn(100)
	fmt.Println("Random Number:", randomNumber)
}
