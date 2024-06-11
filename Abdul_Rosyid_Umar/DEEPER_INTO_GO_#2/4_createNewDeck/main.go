package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// membuat tipe data Deck dari tipe data array String
type Deck []string

// membuat fungsi yang mengembalikan deck atau kartu
func NewDeck() Deck {
	//membuat variabel untuk menampung kartu
	cards := Deck{}
	//membuat variabel untuk menampung komponen kartu
	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	//menyatukan komponen kartu menggunakan perulangan
	for _, suit := range cardSuit {
		//untuk setiap value yang ada di cardsuit
		for _, value := range cardValue {
			//untuk setiap value yang ada di cardvalue
			cards = append(cards, value+" of "+suit)
			//tambahkan value dengan suit lalu ditengahi of untuk setiap value yang ada
		}
	}
	return cards
}

// Deck to string
// Fungsi untuk mengonversi deck menjadi string
func (d Deck) tostring() string {
	//fungsi untuk join slice menjadi string dengan dipisahkan oleh ,
	return strings.Join([]string(d), ", ")
}

// fungsi untuk mengacak kartu
func (d Deck) Shuffle() {
	//digunakan agar setiap kali program dijalankan maka hasilnya berbeda beda
	rand.Seed(time.Now().UnixNano())
	//mengacak elemen berdasarkan panjangnya
	rand.Shuffle(len(d), func(i, j int) {
		//pertukaran elemennya di sini
		d[i], d[j] = d[j], d[i]
	})
}

func randomnumber() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) // Generates a random number between 0 and 99
	fmt.Println("Random number:", randomNumber)
}

func main() {
	deck := NewDeck()
	for i, card := range deck {
		fmt.Println(i+1, card)
	}

	fmt.Println(deck.tostring())
	deck.Shuffle()
	fmt.Println("after shuffle: ", deck)
	randomnumber()
}

// package main

// import "fmt"

// // Definisi tipe baru `Deck`
// type Deck []string

// // Fungsi untuk membuat deck baru
// func NewDeck() Deck {
// 	cards := Deck{}
// 	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
// 	cardValues := []string{"Ace", "Two", "Three", "Four"}

// 	for _, suit := range cardSuits {
// 		for _, value := range cardValues {
// 			cards = append(cards, value+" of "+suit)
// 		}
// 	}
// 	return cards
// }

// func main() {
// 	deck := NewDeck()
// 	for i, card := range deck {
// 		fmt.Println(i+1, card)
// 	}
// }
