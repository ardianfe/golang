package main

import "fmt"

// WhatsMap
// Map adalah koleksi pasangan kunci-nilai (key-value pairs) yang digunakan untuk menyimpan data berdasarkan kunci unik.

func main() {
	// 1. Defining and Initializing a Map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 27,
	}
	fmt.Println("Initial map:", ages)

	// 2. Manipulating Map
	// Adding an element
	ages["Dave"] = 22
	fmt.Println("After adding Dave:", ages)

	// Updating an element
	ages["Alice"] = 26
	fmt.Println("After updating Alice:", ages)

	// Deleting an element
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// Checking if a key exists
	if age, exists := ages["Carol"]; exists {
		fmt.Printf("Carol's age: %d\n", age)
	} else {
		fmt.Println("Carol not found")
	}

	// Iterating Over Map
	fmt.Println("Iterating over map:")
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// Difference Between Maps and Structs
	// Map: dynamic, keys can be added or removed
	// Struct: static, fields are fixed

	// Example with struct
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Eve", Age: 29}
	fmt.Println("Initial person struct:", person)

	// Updating a struct field
	person.Age = 30
	fmt.Println("After updating person struct:", person)

	// Test
	fmt.Println("Running tests:")
	testMapOperations()
}

// Test function to demonstrate map operations
func testMapOperations() {
	// Initializing a map
	testMap := make(map[string]int)
	testMap["One"] = 1
	testMap["Two"] = 2

	// Test addition
	testMap["Three"] = 3
	if testMap["Three"] != 3 {
		fmt.Println("Test failed: addition")
	} else {
		fmt.Println("Test passed: addition")
	}

	// Test update
	testMap["Two"] = 22
	if testMap["Two"] != 22 {
		fmt.Println("Test failed: update")
	} else {
		fmt.Println("Test passed: update")
	}

	// Test deletion
	delete(testMap, "One")
	if _, exists := testMap["One"]; exists {
		fmt.Println("Test failed: deletion")
	} else {
		fmt.Println("Test passed: deletion")
	}
}
