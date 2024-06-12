package main

import "fmt"

// Defining Structs

// Person struct defines basic personal information
type Person struct {
	Name string
	Age  int
}

// Employee struct embeds Person and adds employee-specific information
type Employee struct {
	Person
	EmployeeID string
	Salary     float64
}

// Department struct contains a list of Employees
type Department struct {
	Name      string
	Employees []*Employee
}

// Receiver function to print Employee details
func (e *Employee) PrintDetails() {
	fmt.Printf("Employee Name: %s, Age: %d, EmployeeID: %s, Salary: %.2f\n", e.Name, e.Age, e.EmployeeID, e.Salary)
}

// Function to add an Employee to a Department (Pass by Reference)
func (d *Department) AddEmployee(e *Employee) {
	d.Employees = append(d.Employees, e)
}

// Function to update Employee salary (Pass by Reference)
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

// Function to calculate total salary in a Department
func (d Department) TotalSalary() float64 {
	total := 0.0
	for _, e := range d.Employees {
		total += e.Salary
	}
	return total
}

// Main function
func main() {
	// Creating Employees
	employee1 := &Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E001",
		Salary:     5000.0,
	}
	employee2 := &Employee{
		Person:     Person{Name: "Bob", Age: 40},
		EmployeeID: "E002",
		Salary:     6000.0,
	}

	// Printing initial details
	fmt.Println("Initial Employee Details:")
	employee1.PrintDetails()
	employee2.PrintDetails()

	// Creating Department
	department := Department{Name: "Engineering"}

	// Adding Employees to Department (Pass by Reference)
	department.AddEmployee(employee1)
	department.AddEmployee(employee2)

	// Printing Department details
	fmt.Println("\nDepartment Details:")
	fmt.Printf("Department Name: %s\n", department.Name)
	for _, emp := range department.Employees {
		emp.PrintDetails()
	}

	// Updating Employee salary (Pass by Reference)
	fmt.Println("\nUpdating Employee Salary:")
	employee1.UpdateSalary(5500.0)
	employee1.PrintDetails()

	// Printing Department details again to reflect updated salary
	fmt.Println("\nUpdated Department Details:")
	fmt.Printf("Department Name: %s\n", department.Name)
	for _, emp := range department.Employees {
		emp.PrintDetails()
	}

	// Calculating total salary in the Department
	fmt.Printf("\nTotal Salary in %s Department: %.2f\n", department.Name, department.TotalSalary())

	// Demonstrating Pass by Value (Original employee2 remains unchanged)
	updateEmployeeAge(*employee2, 45)
	fmt.Println("\nAfter attempting to update age (Pass by Value):")
	employee2.PrintDetails()

	// Demonstrating Pass by Reference (Original employee2 gets updated)
	updateEmployeeAgeReference(employee2, 45)
	fmt.Println("\nAfter updating age (Pass by Reference):")
	employee2.PrintDetails()
}

// Function to update Employee age (Pass by Value)
func updateEmployeeAge(e Employee, newAge int) {
	e.Age = newAge
	fmt.Println("Inside updateEmployeeAge (Pass by Value):")
	e.PrintDetails()
}

// Function to update Employee age (Pass by Reference)
func updateEmployeeAgeReference(e *Employee, newAge int) {
	e.Age = newAge
	fmt.Println("Inside updateEmployeeAgeReference (Pass by Reference):")
	e.PrintDetails()
}
