package main

import "fmt"

// Proses assertion adalah proses memeriksa apakah suatu elemen ada didalam slice. biasanya dilakukan dengan cara perulangan atau iterasi melalui setiap elemen slice dan memebandingan dengan nilai yang dicari. misal nilai tersebut ditemukan maka dapat disimpukan bahwa element ada dalam slice

// mengecek nilai dalam slice menggunakan boolean
func assertElementInSlice(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func assertion(slice []string, isi string) bool {
	for _, v := range slice {
		if v == isi {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	target := 7

	// jadi kita membuat kondisi untuk mengecek apabila element tersebut termasuk dalam slice maka akan menghasilkan true dan apa bila gaada pada slice maka false

	if assertElementInSlice(slice, target) {
		fmt.Printf("Elemen %d ditemukan dalam slice \n",target)
	} else {
		fmt.Printf("Element %d tidak dimukan dalam slice \n",target)
	}

	// ==== string ====

	sliceString := []string{"Deni", "Yudi", "Saputra"}

	targetString := "Saputraw"

	if assertion(sliceString, targetString) {
		fmt.Printf("%s terdapat pada slice", targetString)
	} else {
		fmt.Printf("%s tidak terdapat pada slice", targetString)
	}
}
