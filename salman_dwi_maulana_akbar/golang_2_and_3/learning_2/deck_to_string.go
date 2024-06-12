package main

import (
	"fmt"
	"strings"
)

type Card struct {
    Suit  string
    Value string
}

type Deck []Card

func (d Deck) ToString() string {
    var cards []string
    for _, card := range d {
        cards = append(cards, card.Value+" of "+card.Suit)
    }
    return strings.Join(cards, ", ")
}

func NewDeck() Deck {
    suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
    values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
    var deck Deck

    for _, suit := range suits {
        for _, value := range values {
            deck = append(deck, Card{Suit: suit, Value: value})
        }
    }
    return deck
}

func main() {
    deck := NewDeck()
    fmt.Println(deck.ToString())
}
