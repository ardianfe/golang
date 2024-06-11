package main

import "fmt"

func main() {
	str := "hello"
	byteSlice := []byte(str)

	fmt.Println(byteSlice) // Output: [104 101 108 108 111]
}
