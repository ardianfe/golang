package main

import "fmt"

func main() {
	m1 := map[string]int{"one": 1, "two": 2}
	m2 := m1       // m2 merujuk ke map yang sama dengan m1
	m2["one"] = 10 // Mengubah m2 juga mempengaruhi m1

	fmt.Println(m1) // Output: map[one:10 two:2]
	fmt.Println(m2) // Output: map[one:10 two:2]
}
