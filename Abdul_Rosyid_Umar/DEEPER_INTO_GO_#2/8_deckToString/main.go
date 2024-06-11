package main

import (
	"fmt"
	"strings"
)

type Deck []string

func NewDeck() Deck {
	cards := Deck{"Ace of Spades", "Two of Diamonds"}
	return cards
}

// Fungsi untuk mengonversi deck menjadi string
func (d Deck) ToString() string {
	return strings.Join([]string(d), ", ")
}

func main() {
	deck := NewDeck()
	fmt.Println(deck.ToString()) // Output: Ace of Spades, Two of Diamonds
}
