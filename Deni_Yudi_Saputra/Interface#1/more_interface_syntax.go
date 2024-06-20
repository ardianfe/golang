package main

import "fmt"

// Definisikan sebuah interface dengan satu metode
type Sapaan interface {
	Sapa() string
}

// Definisikan sebuah struct yang akan mengimplementasikan interface
type Manusia struct {
	Nama string
}

// Implementasikan metode Sapa() dari interface Sapaan pada struct Manusia
func (m Manusia) Sapa() string {
	return "Halo, nama saya " + m.Nama
}

// Fungsi untuk menerima argumen yang bertipe Sapaan
func SapaSemua(s Sapaan) {
	fmt.Println(s.Sapa())
}

func moreInterface() {
	// Buat objek Manusia dengan nama Deni
	deni := Manusia{"Deni"}

	// Panggil fungsi SapaSemua dengan argumen objek Manusia
	SapaSemua(deni)

	// Deklarasi variabel bertipe interface Sapaan
	var sapaan Sapaan

	// Inisialisasi variabel sapaan dengan objek Manusia
	sapaan = Manusia{"Budi"}

	// Panggil metode Sapa() dari variabel sapaan
	fmt.Println(sapaan.Sapa())
}

func main() {
	moreInterface()
}
