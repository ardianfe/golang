package main

import (
	"fmt"
	"strings"
)

func main() {
    words := []string{"Go", "is", "fun"}
    sentence := strings.Join(words, " ")
    fmt.Println(sentence)
}
