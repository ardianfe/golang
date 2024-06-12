package main

import "fmt"

func main() {

	var a int = 10
	var p *int // deklarasi pointer p

	p = &a //p menyimpan alamat variable a

	fmt.Println("Nilai a :", a) //nilai a
	fmt.Println("Alamat a :", &a) // alamat a (alamat memori) 0x123123123
	fmt.Println("Nilai p :", p) // alamat a
	fmt.Println("Nilai *p :", *p) //nilai *p = 10 (alamat a)

	*p = 20 //mengubah nilai menjadi pointer p
	fmt.Println("Nilai a setelah diubah melalui pointer:", a) // hasil setelah dirubah menggunakan pointer 

}

// Pointer adalah variable yang menyimpan alamat memori variable lain
// & digunakan untuk mendapakan memori dari suatu variable 
// * digunakan untuk mengakses nilai dari alamat yang ditunjuk oleh pointer 