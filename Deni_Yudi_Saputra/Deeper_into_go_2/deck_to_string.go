package main

import (
	"fmt"
	"strings"
)

type CardString struct {
	Suit  string
	Value string
}

type DeckString []CardString

// method string mengubah deck menjadi representasi string
func (d DeckString) String() string {
	var sb strings.Builder

	for _, card := range d {
		sb.WriteString(fmt.Sprintf("%s dari %s\n", card.Value, card.Suit))
	}

	return sb.String()
}
func main() {
	// Membuat contoh dek kartu
	deck := DeckString{
		{Suit: "Hearts", Value: "Ace"},
		{Suit: "Diamonds", Value: "2"},
		{Suit: "Clubs", Value: "3"},
		{Suit: "Spades", Value: "4"},
	}

	// Mengubah dek menjadi string dan mencetaknya
	fmt.Println(deck.String())
}

// PERBEDAAN MENGGUNAKAN STRING DAN +

// untuk menggunakan string.Builder ialah lebih efisiensi waktu dan alokasi memori lebih sedikit menjadikan lebih cepat dari pada +
// untuk "+=" go akan mengalokasikan string baru dan menyalin konten lama serta konten baru kedalamnya. itulah yang menyebabkan memori yang besar dibanding dengan strings
// contoh menggunakan string dan +

// ===STRINGS ====
// func (d Deck) String() string {
//     var result string
//     for _, card := range d {
//         result += fmt.Sprintf("%s of %s\n", card.Value, card.Suit)
//     }
//     return result
// }

// ===== "+="======
// func (d Deck) String() string {
//     var result string
//     for _, card := range d {
//         result += fmt.Sprintf("%s of %s\n", card.Value, card.Suit)
//     }
//     return result
// }

// Sebenarnya sama saja cuman berbeda pada efisiensi dan penggunaan memori
