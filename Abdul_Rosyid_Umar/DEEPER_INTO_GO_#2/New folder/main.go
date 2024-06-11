package main

import "fmt"

// Fungsi yang mengembalikan dua nilai
func divide(a, b int) (int, int, int) {
	hasilbagi := a / b
	sisa := a % b
	jikadikali := a * b
	return hasilbagi, sisa, jikadikali
}

func main() {
	q, r, m := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d, multiple: %d\n", q, r, m)
}
