package main

import "fmt"

type Person1 struct {
	FirstName, LastName string
}

type Company1 struct {
	Name string
}

func (p Person1) GetName() string {
	return p.FirstName + " " + p.LastName
}
func (c Company1) GetName() string {
	return c.Name
}

type Named interface {
	GetName() string // bisa membuat satu method, tapi method tersebut mempunyai receiver berbeda serta hasil yang berbeda
}

func printName(n Named) {
	fmt.Println("Name", n.GetName())
}

func main() {

	person := Person1{"Dw", "ati"}
	company := Company1{"Qualita"}

	printName(person)
	printName(company)
}
