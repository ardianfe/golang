package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []string

func NewDeck() Deck {
	cards := Deck{"Ace of Spades", "Two of Diamonds", "Three of Hearts", "Four of Clubs"}
	return cards
}

// fungsi untuk shuffle deck
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func main() {
	deck := NewDeck()
	fmt.Println("Before shuffle:", deck)
	deck.Shuffle()
	fmt.Println("After shuffle:", deck)
}
