package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "World"}
	sentence := strings.Join(words, " ")

	fmt.Println(sentence) // Output: Hello World
}
