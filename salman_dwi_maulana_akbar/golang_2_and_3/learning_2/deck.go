package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Cards []string
}

func NewDeck() Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	var cards []string

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" "+suit)
		}
	}

	return Deck{Cards: cards}
}

func (d *Deck) ToString() string {
	return strings.Join(d.Cards, ", ")
}

func (d *Deck) AddCard(card string) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck) RemoveCard(index int) {
    if index >= 0 && index < len(d.Cards) {
        d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
    }
}

func (d *Deck) AccessCard(index int) string {
	if index >= 0 && index < len(d.Cards) {
		return d.Cards[index]
	}
	return ""
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Cut(position int) (Deck, Deck) {
	if position >= 0 && position < len(d.Cards) {
		left := Deck{Cards: d.Cards[:position]}
		right := Deck{Cards: d.Cards[position:]}
		return left, right
	}
	return Deck{}, Deck{}
}

func (d *Deck) Merge(other Deck) {
	d.Cards = append(d.Cards, other.Cards...)
}

func main() {
	// Membuat deck baru
	fmt.Println("Membuat deck baru...")
	deck := NewDeck()
	fmt.Println("Deck baru berhasil dibuat.")

	// Mencetak isi deck
	fmt.Println("\nIsi deck:")
	fmt.Println(deck.ToString())

	// Menambahkan kartu ke deck
	fmt.Println("\nMenambahkan kartu ke deck...")
	deck.AddCard("As Hati")
	deck.AddCard("Dua Sekop")
	fmt.Println("Kartu berhasil ditambahkan.")

	// Mencetak isi deck setelah penambahan kartu
	fmt.Println("\nIsi deck setelah penambahan kartu:")
	fmt.Println(deck.ToString())

	// Menghapus kartu dari deck
	fmt.Println("\nMenghapus kartu dari deck...")
	deck.RemoveCard(0) // Hapus kartu pertama
	fmt.Println("Kartu berhasil dihapus.")

	// Mencetak isi deck setelah penghapusan kartu
	fmt.Println("\nIsi deck setelah penghapusan kartu:")
	fmt.Println(deck.ToString())

	// Mengakses kartu dalam deck
	fmt.Println("\nMengakses kartu dalam deck...")
	card := deck.AccessCard(2) // Akses kartu ketiga
	fmt.Println("Kartu yang diakses:", card)

	// Mengacak deck
	fmt.Println("\nMengacak deck...")
	deck.Shuffle()
	fmt.Println("Deck berhasil diacak.")

	// Mencetak isi deck setelah pengacakan
	fmt.Println("\nIsi deck setelah pengacakan:")
	fmt.Println(deck.ToString())

	// Memotong deck
	fmt.Println("\nMemotong deck...")
	left, right := deck.Cut(5) // Memotong deck pada posisi ke-5
	fmt.Println("Deck kiri setelah pemotongan:")
	fmt.Println(left.ToString())
	fmt.Println("\nDeck kanan setelah pemotongan:")
	fmt.Println(right.ToString())

	// Menggabungkan deck
	fmt.Println("\nMenggabungkan deck...")
	left.Merge(right)
	fmt.Println("Deck berhasil digabungkan.")

	// Mencetak isi deck setelah penggabungan
	fmt.Println("\nIsi deck setelah penggabungan:")
	fmt.Println(left.ToString())

	fmt.Println("\nMembagikan kartu kepada dua pemain...")
	player1, remainingDeck := deck.Cut(5)
	player2, remainingDeck := remainingDeck.Cut(5)
	fmt.Println("Kartu pemain 1:", player1.ToString())
	fmt.Println("Kartu pemain 2:", player2.ToString())
	fmt.Println("Sisa deck:", remainingDeck.ToString())

	// Pemain mengembalikan kartu mereka ke deck
	fmt.Println("\nPemain mengembalikan kartu mereka ke deck...")
	remainingDeck.Merge(player1)
	remainingDeck.Merge(player2)
	fmt.Println("Isi deck setelah penggabungan kartu pemain:", remainingDeck.ToString())

	// Mengacak deck lagi setelah pengembalian kartu
	remainingDeck.Shuffle()
	fmt.Println("Deck berhasil diacak kembali.")
	fmt.Println("Isi deck setelah diacak kembali:", remainingDeck.ToString())
}
