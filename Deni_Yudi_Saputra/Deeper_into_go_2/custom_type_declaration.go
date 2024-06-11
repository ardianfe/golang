package main

import "fmt"

type Jeneng string

func gabungNama(Jeneng string) {
	var nama = "Deni Yudi "

	var dadiSiji = nama + Jeneng

	fmt.Println("Halo", dadiSiji)
}

type Jumlah int

var j Jumlah

func Penjumlahan(tambah int) {

	var jo = 20
	var tambahan = jo + tambah

	fmt.Println(tambahan)

}

func main() {

	Penjumlahan(20)
	gabungNama("Saputra")
}

// // mendeklarasikan tipe custom berdasarkan string
// type namaDepan string

// type namaLengkap string

// // tambah method pada custom type declaration

// func (nama namaLengkap) Cetak() string {
// 	return string(nama)
// }

// func main() {
// 	var namaSaya namaDepan = "Deni Yudi Saputra"

// 	fmt.Println("nama saya : ", namaSaya)

// 	var jeneng namaLengkap = "Jono-Joni"

// 	fmt.Println("Nama lengkap saya : ", jeneng.Cetak())

// }

// type NamaDepan string

// type Pesan string // type alias

// type Koordinat struct {
// 	Latitude  float64
// 	Longitude float64
// }

// func hitungJarak(a Koordinat, b Koordinat) float64 {
// 	// jarak := a.Latitude + b.Latitude
// 	return 100
// }

// func main() {
// 	pesan := Pesan("Halo, gimana?")
// 	fmt.Println(pesan)

// 	var malang Koordinat = Koordinat{7.983333, 112.616667}
// 	fmt.Println("Koordinat Malang:", malang)

// 	var jakarta Koordinat = Koordinat{-6.214000, 106.845000}
// 	fmt.Println("Jarak ke Jakarta:", hitungJarak(malang, jakarta))
// }
