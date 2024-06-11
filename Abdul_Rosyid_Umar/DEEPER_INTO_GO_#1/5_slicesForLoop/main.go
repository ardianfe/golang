package main

import "fmt"

func main() {
	//declarasi slice bernama numbers
	var numbers = []int{1, 2, 3, 4, 5}

	//menggunakan for untuk pengulangan dan print setiap number yang ada di slice
	fmt.Println("Numbers:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	//menggunakna for range untuk pengulangan dan print setiap number yang ada di slice per indexnya
	fmt.Println("Using range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
