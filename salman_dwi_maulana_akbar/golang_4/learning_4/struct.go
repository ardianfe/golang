package main

import "fmt"

// 1. Defining Struct
type Person struct {
	Name string
	Age  int
}

// 3. Embedding Structs
type Employee struct {
	Person
	EmployeeID string
}

// Pass by Value
func updatePersonValue(p Person) {
	p.Age = 40
	fmt.Println("Inside updatePersonValue:", p)
}

// Structs with Pointers (Pass by Reference)
func updatePersonReference(p *Person) {
	p.Age = 50
	fmt.Println("Inside updatePersonReference:", *p)
}


	// Structs with Receiver Functions
	func (e Employee) PrintDetails() {
		fmt.Printf("Employee Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
	}	

// 2. Declaring Struct and Updating Struct Value
func main() {
	// Declaring a struct
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Initial person:", person)

	// Updating struct value
	person.Age = 30
	fmt.Println("Updated person:", person)

	// 3. Embedding Structs
	
	employee := Employee{
		Person:     Person{Name: "Bob", Age: 35},
		EmployeeID: "E123",
	}
	fmt.Println("Employee:", employee)

	// 4. Structs with Receiver Functions
	employee.PrintDetails()
	
	// 5. Pass by Value
	updatePersonValue(person)
	fmt.Println("After updatePersonValue:", person)

	// 6. Structs with Pointers (Pass by Reference)
	updatePersonReference(&person)
	fmt.Println("After updatePersonReference:", person)

	// 7. Pointer Operations
	ptr := &person
	fmt.Println("Person via pointer:", *ptr)
	ptr.Age = 35
	fmt.Println("Person after pointer modification:", person)

	// 8. Pointer Shortcuts
	ptr = new(Person)
	ptr.Name = "Charlie"
	ptr.Age = 28
	fmt.Println("New person via pointer shortcut:", *ptr)

	// 9. Gotchas with Pointers
	var uninitializedPtr *Person
	if uninitializedPtr == nil {
		fmt.Println("Uninitialized pointer is nil")
	}

	// 10. Reference vs Value Type
	var intValue int = 10
	var intPtr *int = &intValue
	fmt.Println("Initial intValue:", intValue)
	*intPtr = 20
	fmt.Println("Updated intValue via pointer:", intValue)
}



