package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	//mengambil nilai slice dari index 1 sampai index sebelum 4
	subset := number[1:4]
	fmt.Println(subset)

	//Byte slice
	str := "hellooow"
	//mengconversi helloow menjadi nomor binary
	bytstr := []byte(str)
	fmt.Println(bytstr)
}
