package main

import (
	"fmt"
	"strings"
)

func main() {
	cards := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	joined := strings.Join(cards, ", ")
	fmt.Println("Deck:", joined)
}
