package main

import "fmt"

func main() {
	x := 10
	p := &x // p adalah pointer yang menyimpan alamat memori dari x

	fmt.Println(*p) // Output: 10, *p mengakses nilai x melalui pointer
	*p = 20         // Mengubah nilai x melalui pointer
	fmt.Println(x)  // Output: 20
}
