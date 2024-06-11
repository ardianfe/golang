package main

import (
	"fmt"
	"strings"
)

// jadi join slice adalah kita harus menggunakan strings supaya bisa join slice yang bertipe data string dengan menambahkan Join

func main() {

	// slice string
	kata := []string{"saya", "deni", "slebew"}

	fmt.Println(kata)

	// gabung slice untuk memisahkan kata
	hasil := strings.Join(kata, "_")
	fmt.Println(hasil)

	// angka := []int{1, 2, 3, 4, 5} // gabisa soalnya value dari slice  harus string

	// fmt.Println(angka)

	// hasil2 := strings.Join(angka,"+")

	// fmt.Println(hasil2)
}
