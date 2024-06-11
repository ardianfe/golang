package main

import "fmt"

func main() {
	// var name string

	// name = "ini contoh var"
	// fmt.Println(name)

	// name = "ini contoh var 2"
	// fmt.Println(name)

	// golang tahu tipe data nya asal setelah var kita mendefinisikan variablenya

	// var name = "ini contoh var"
	// fmt.Println(name)

	// name = "ini contoh var 2"
	// fmt.Println(name)

	// kita juga bisa mendeklarasikan var dengan mengunakan ":" sebelum tanda "=" di var kita
	name := "ini contoh var"
	fmt.Println(name)

	name = "ini contoh var 2"
	fmt.Println(name)

	// di golang kita juga bisa mendeklarasikan multiple variable seperti contoh dibawah ini

	var (
		multi1 = "var 1"
		multi2 = "var 2"
		multi3 = "var 3"
	)

	fmt.Println(multi1)
	fmt.Println(multi2)
	fmt.Println(multi3)
}
