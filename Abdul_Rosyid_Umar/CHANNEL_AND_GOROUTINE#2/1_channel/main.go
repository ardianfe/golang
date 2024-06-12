package main

import "fmt"

func main() {
	ch := make(chan int) // mmebuat channel

	func() {
		ch <- 1 // memasukan data ke channel
		ch <- 2
		ch <- 3
		close(ch) // menutup channel
	}()

	for val := range ch {
		fmt.Println("received: ", val) // hasil output akan eror karena tidak memakai goroutine yang membuat channel tidak memiliki penerima
	}
}
