package main

import "fmt"

func main() {
	var x int = 10  // x adalah sebuah variabel bertipe int dan bernilai 10
	var p *int = &x // p adalah pointer yang menyimpan alamat memori dari x

	fmt.Println("Nilai x:", x)            // Mencetak nilai x (10)
	fmt.Println("Alamat x:", p)           // Mencetak alamat memori dari x (misalnya 0x140000b2c08)
	fmt.Println("Nilai di alamat p:", *p) // Mencetak nilai yang disimpan di alamat p (10)

	*p = 20                         // Mengubah nilai yang disimpan di alamat p, ini berarti mengubah nilai x menjadi 20
	fmt.Println("Nilai baru x:", x) // Mencetak nilai x yang baru (20)
}
