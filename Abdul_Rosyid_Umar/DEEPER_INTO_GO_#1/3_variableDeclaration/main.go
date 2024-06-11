package main

import "fmt"

//variable declaration

func main() {
	//menggunakan var
	//var a int

	//declarasi sekaligus inisialisasi
	var a int = 10

	//declarasi tanpa tipe data (tipe ditentukan secara otomatis)
	var b = 20

	//declarasi shorthand (hanya dalam fungsi)
	c := 30

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
