package main

import "fmt"

func main() {

	var slice []int // deklarasi slice kosong

	slice = []int{1, 2, 3, 4, 5}

	subSlice1 := slice[1:3]
	fmt.Println("Sub-slice 1 :", subSlice1) //2,3

	subSlice2 := slice[:2]
	fmt.Println("Sub-slice 2 :", subSlice2) //1,2

	subSlice3 := slice[3:]
	fmt.Println("Sub-slice 3 :", subSlice3) //4,5

	// buat slice baru menggunakan make untuk mementukan panjang dan kapastias dari slice

	newSlice := make([]int, 1, 6)

	fmt.Println("New Slice :", newSlice)

	// ketika sudah membuat new slice menggunakan make maka menambahkan isian slice dengan menggunakan append

	newSlice = append(newSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("New Slice setelah append :", newSlice)

	newSlice1 := newSlice[2:7]
	fmt.Println("hasil dari :", newSlice1)

	// perulangan dengan for range

	for i, v := range newSlice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
