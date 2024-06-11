// for
// dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
// salah satu fitur perulangan adalah for loops

package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
	fmt.Println("selesai")

	// for dengan statement
	// dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
	// init statement, yaitu statement sebelum for di eksekusi
	// post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
	// perulangan yang lebih simple

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	// for range
	// for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// data collection contohnya array, slice, dan map

	// sebelum memakai for range
	name := []string{"anjay", "mabar", "bro"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	// setelah memakai for range
	for index, name := range name {
		fmt.Println("index", index, "=", name)
	}
	// jika tidak membutuhkan indexnya
	for _, name := range name {
		fmt.Println(name)
	}

	// break & continue
	// break & continue adalah kata kunci yang bisa digunakan dalam perulangan
	// break digunakan untuk menghentikan seluruh perulangan
	// continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

	for a := 0; a < 10; a++ {
		if a == 5 {
			break
		}
		fmt.Println("perulangan ke", a)
	}
	// perulangan yang hanya bilangan ganjil atau genap tergantung yang buat codenya atau yang bisa disebut continue
	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", a)
	}

}
