package main

import "fmt"

type Book struct {
	Judul string
	Harga float64
}

//fungsi untuk mengubah harga menggunakan parameter book (pass by value)
func ubahHarga(b Book) {
	b.Harga = 12000
	fmt.Println("Harga Buku:", b.Harga)
}

func main() {
	buku1 := Book{"Bumi", 5000}
	ubahHarga(buku1)
	fmt.Println("Harga Buku:", buku1.Harga) //price tetap 5000
}
