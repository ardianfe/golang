package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

type Company struct {
	Name string
}

func (p Person) GetName() string {
	return p.FirstName + " " + p.LastName
}
func (c Company) GetName() string {
	return c.Name
}

func main() {
	person := Person{"Deni", "Yudi"}
	company := Company{"Qualita"}

	fmt.Println("Person name :",person.GetName())
	fmt.Println("Company name :",company.GetName())
}
