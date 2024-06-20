package main

import (
	"fmt"
	"time"
)

func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	go func() {
		fmt.Println("Ini adalah goroutine anonim")
	}()

	time.Sleep(1 * time.Second) // Menunggu goroutine selesai

	// return value literal
	sum := func(a int, b int) int {
		return a + b
	}

	result := sum(3, 4)
	fmt.Println("Hasil penjumlahan:", result) // Output: Hasil penjumlahan: 7

	//pada fungsi
	result = operate(5, 3, func(x int, y int) int {
		return x - y
	})

	fmt.Println("Hasil operasi:", result) // Output: Hasil operasi: 2

	//Loop pada literal

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
