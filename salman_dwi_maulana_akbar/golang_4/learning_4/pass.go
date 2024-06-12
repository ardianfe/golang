package main

import "fmt"

// Definisi struct Person
type Person struct {
    Name string
    Age  int
}

// Pass by reference: Mengubah umur pada referensi struct
func updatePersonReference(p *Person) {
    p.Age = 30
    fmt.Println("Inside function (updatePersonReference):", *p) // Age: 30
}

// Pass by value: Mencoba mengubah umur pada salinan struct
func updatePersonValue(p Person) {
    p.Age = 20
    fmt.Println("Inside function (updatePersonValue):", p) // Age: 30
}

func main() {
    // Inisialisasi struct Person
    person := Person{Name: "Alice", Age: 25}
				person1 := Person{Name: "Bob", Age: 30}

    fmt.Println("Before function calls Alice:", person) // Age: 25
				fmt.Println("Before function calls Bob:", person1) // Age: 30

    // Pass by reference
    updatePersonReference(&person)
    fmt.Println("After updatePersonReference call:", person) // Age: 30

    // Pass by value
    updatePersonValue(person1)
    fmt.Println("After updatePersonValue call:", person1) // Age: 30
}
