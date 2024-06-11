package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Fungsi dengan receiver `Person`
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "John", Age: 90}
	p.Greet()
}
