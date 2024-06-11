package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deklarasi struct
type Kartu struct {
	Jenis, Nilai string
}

// Deklarasi type khusus
type Deck []Kartu

// fungsi penerima cetak
func (d Deck) Cetak() {
	for _, v := range d {
		fmt.Println(v.Jenis, v.Nilai)
	}
}

// fungsi penerima kestring
func (d Deck) keString() string {
	var kartuStrings []string
	for _, kartu := range d {
		kartuStrings = append(kartuStrings, kartu.Jenis+" "+kartu.Nilai)
	}
	return strings.Join(kartuStrings, " ")
}

// fungsi buat deck
func BuatDeck() Deck {
	Jenis := []string{"Hati", "Wajik", "Sekop", "Keriting"}
	Nilai := []string{"AS", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "k"}
	var d Deck

	for _, j := range Jenis {
		for _, n := range Nilai {
			d = append(d, Kartu{Jenis: j, Nilai: n})
		}
	}
	return d
}

// fungsi yang membagi deck menjadi dua bagian
func BagiDek(d Deck) (Deck, Deck) {
	bagi := len(d) / 2
	return d[:bagi], d[bagi:]
}

// menggabung irisan string
func (d Deck) Gabungkan() string {
	var kartuStrings []string
	for _, kartu := range d {
		kartuStrings = append(kartuStrings, kartu.Jenis+" "+kartu.Nilai)
	}
	return strings.Join(kartuStrings, ", ")
}

// func (d Deck) SimpanKeFile(nama string) error {
// 	return ioutil.WriteFile(nama, []byte(d.keString()), 0644)

// }

// Fungsi menyimpan file
func (d Deck) SimpanKeFile(nama string) {
	err := ioutil.WriteFile(nama, []byte(d.Gabungkan()), 0644)
	if err != nil {
		fmt.Println("Error dalam menyimpan file: ", err)
	}
}

// func bacaDariFile(namaFile string) (Deck, error) {
// 	bs, err := ioutil.ReadFile(namaFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("error dalam membaca data: %v", err)
// 	}
// 	s := strings.Split(string(bs), ", ")
// 	var dek Deck
// 	for _, kartuString := range s {
// 		kartuData := strings.Split(kartuString, " ")
// 		dek = append(dek, Kartu{Jenis: kartuData[0], Nilai: kartuData[1]})
// 	}
// 	return dek, nil
// }

// fungsi baca dari file
func bacaDariFile(namaFile string) Deck {
	bs, err := ioutil.ReadFile(namaFile)
	if err != nil {
		fmt.Println("Error dalam membaca file: ", err)
	}
	s := strings.Split(string(bs), ", ")
	var dek Deck
	for _, kartuString := range s {
		kartuData := strings.Split(kartuString, " ")
		dek = append(dek, Kartu{Jenis: kartuData[0], Nilai: kartuData[1]})
	}
	return dek
}

// fungsi acak
func (d Deck) acak() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func main() {
	// membuat deck baru
	dek := BuatDeck()
	dek.Cetak()

	// deck menjadi string
	fmt.Println("dek diubah jadi string", dek.keString())
	fmt.Println("dek diubah jadi string dengan dibatasi koma", dek.Gabungkan())

	// membagi deck
	bagian1, bagian2 := BagiDek(dek)
	fmt.Printf("ini bagian 1\n")
	bagian1.Cetak()
	fmt.Printf("ini bagian 2\n")
	bagian2.Cetak()

	// mengacak Deck dan menyimpannya ke file
	bagian1.acak()
	bagian1.SimpanKeFile("example.docx")

	// membaca dari file
	bacaDariFile("example.docx")
}
