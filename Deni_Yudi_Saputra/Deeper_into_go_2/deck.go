package main

import "fmt"

type Card struct{
	Suit string
	Value string
}

var suits = []string {"Keriting","Wajik","Hati","Sekop"}
var values = []string {"2","3","4","5","6","7","8","9","10","J","Q","K","A"}

// tambah deck

func Deck() []Card {
	var deck  []Card

	for _, suit := range suits{
		for _, value := range values{
			// penggunaan append pada range pada go diperlukan karena kita perlu menambahkan elemen baru ke dalam slice pada setiap iterasi
			 deck = append(deck,Card{Suit: suit,Value: value})
		}
	}

	return deck
}

func PrintDeck(deck []Card){
	for _, card := range deck{
		fmt.Println(card.Value,"dari",card.Suit)
	}
}

func main()  {
	PrintDeck(Deck())
}





// ====== MENAMPILKAN DECK NYA SAJA =========

// type Card struct {
// 	Suit  string
// 	Value string
// }

// var suits = []string{"Hati", "Sekop", "Keriting", "Wajik"}
// var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// func NewDeck() []Card {
// 	var deck []Card

// 	// perulangan range buat menampilkan suits dan values, untuk dimasukkan ke variable deck dengan slice kosong Card 
// 	for _, suit := range suits {
// 		for _, value := range values {
// 			deck = append(deck, Card{Suit: suit, Value: value})
// 		}
// 	}

// 	return deck
// }

// func PrintDeck(deck []Card){
// 	for _, card:= range deck {
// 		fmt.Println(card.Value,"dari",card.Suit )
// 	}
// }

// func main() {
// 	PrintDeck(NewDeck())
// }

//  ===== DECK 1 =====
// type Card struct {
// 	Suit  string
// 	Value string
// }

// func (c Card) String() string {
// 	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
// }

// type Deck []Card

// func NewDeck() Deck {
// 	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
// 	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

// 	var deck Deck

// 	for _, suit := range suits {
// 		for _, value := range values {
// 			deck = append(deck, Card{Suit: suit, Value: value})
// 		}
// 	}

// 	return deck
// }

// // mengocok kartu
// func (d Deck) Shuffle() {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := range d {
// 		j := rand.Intn(i + 1)
// 		d[i], d[j] = d[j], d[i]
// 	}
// }

// // menampilkan kartu
// func (d Deck) Print() {
// 	for _, card := range d {
// 		fmt.Println(card)
// 	}
// }

// func main() {
// 	deck := NewDeck()
// 	deck.Shuffle()
// 	deck.Print()
// }
