package main

import "fmt"

type Penulis struct {
	Nama, Gender string
}

type Buku struct {
	Judul string
	//embedding struct penulis, bisa disebut juga pewarisan atau inheritance di golang
	Penulis
	Jumlah int
	Harga  float64
}

func main() {
	buku1 := Buku{
		Judul: "Bumi",
		Penulis: Penulis{
			Nama:   "Ahmad",
			Gender: "Laki-laki",
		},
		Jumlah: 19,
		Harga:  100000,
	}
	fmt.Printf(buku1.Nama)
}
