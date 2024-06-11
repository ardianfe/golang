package main

import "fmt"

type Buku struct {
	Judul, Penulis string
	Jumlah         int
	Harga          float64
}

// mendeklarasikan variable dari struct Buku yang sudah dibuat
var buku1 Buku

func main() {
	//menginisialisasikan struct
	buku2 := Buku{
		Judul:   "Bumi",
		Penulis: "Tere Liye",
		Jumlah:  100,
		Harga:   25000,
	}

	fmt.Printf(buku2.Judul)
}

//menginisialisasikan struct
