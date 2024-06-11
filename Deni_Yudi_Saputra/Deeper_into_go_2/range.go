package main

import "fmt"

func main() {

	number := []int{1, 22, 30, 40} //index 0-terakhir

	// // perulangan for biasa
	// for i := 0; i < len(number); i++ {
	// 	fmt.Println(number[i])
	// }

	for i, element := range number {
		fmt.Printf("Index %d : %d \n", i, element)
	}


	var buah map[string]int = map[string]int{
		"Apel": 10,
        "Jeruk": 20,
		"Mangga": 30,
		"Pisang": 40,
	}

	for key, val := range buah {
		fmt.Println(key, val)
	}
}
