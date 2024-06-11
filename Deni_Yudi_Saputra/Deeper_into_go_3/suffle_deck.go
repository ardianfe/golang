package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

var suits = []string{"Hati", "Sekop", "Keriting"}
var value = []string{"1", "2", "3", "4"}

var Deck []Card

func NewDeck() []Card {
	// var kartu string
	for _, s := range suits {
		for _, v := range value {
			Deck = append(Deck, Card{Suit: s, Value: v})

		}
	}

	return Deck
}

func PrintDeck(c []Card) {
	for _, v := range c {
		fmt.Println(v)
	}
}

func main() {
	PrintDeck(NewDeck())

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Deck),func (i, j int)  {
		Deck[i],Deck[j]=Deck[j],Deck[i] 
	})

	fmt.Println("Acak :",Deck)
}
