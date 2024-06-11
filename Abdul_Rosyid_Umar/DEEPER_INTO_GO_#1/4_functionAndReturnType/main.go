package main

import "fmt"

//function and return type

func greet() {
	fmt.Println("Hello, World!")
	//fungsi tanpa pengembalian nilai dan tidak menggunakan parameter
}

func greetPerson(name string) {
	fmt.Println("Hello,", name)
	//fungsi tanpa pengembalian nilai dan menggunakan parameter name string
}

func add(a int, b int) int {
	return a + b
	//fungsi mengembalikan nilai int dan menggunakan 2 parameter int a dan b
}

func main() {
	//fungsi yang menggunakan parameter, perlu diisi parameternya
	greet()
	greetPerson("Alice")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
