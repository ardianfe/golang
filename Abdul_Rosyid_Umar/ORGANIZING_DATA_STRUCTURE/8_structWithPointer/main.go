package main

import "fmt"

//mendefinisikna struct
type Book struct {
	Judul string
	Harga float64
}

//membuat fungsi untuk mengubah harga di dalam struct book
func ubahHarga(b *Book, harga float64) {
	b.Harga = harga
	fmt.Printf("Harga buku setelah diubah: %f\n", b.Harga)
}

func main() {
	//menginisialisasi struct buku1
	buku1 := Book{"Bumi", 17000}
	fmt.Printf("harga buku sebelum diubah: %f\n", buku1.Harga)

	//menginisialisasi p untuk menyimpan alamat memori dari buku1
	p := &buku1
	//di sini a juga akan menyimpan alamat memori dari buku1 karena diinisialisasikan dengan p
	a := p
	//karena a juga menyimpan alamat memori buku1 jadi jika b diinisialisasikan dengan a akan menyimpan alamat memori yang sama
	b := a
	//jika kita ubah b maka akan mengubah semuanya karena semuanya menyimpan alamat memori yang sama dari buku1
	ubahHarga(b, 22000.50)
	b.Harga = 22000.50 // ini setara dengan (*b).Harga dimana kita perlu tanda * untuk mengakses apa yang ada di dalam alamatnnya
	fmt.Printf("ini p: %f\n", p.Harga)
	fmt.Printf("ini a: %f\n", a.Harga)
	fmt.Printf("ini b: %f\n", b.Harga)
	fmt.Printf("ini buku1: %f\n", buku1.Harga)
}
