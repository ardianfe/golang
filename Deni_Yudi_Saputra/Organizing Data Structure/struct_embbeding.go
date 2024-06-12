package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type PersonEmbbed struct {
	Name string
	Age  int
	Address
}

func main() {
	person := PersonEmbbed{
		Name: "Deni",
		Age:  32,
		Address: Address{
			Street: "Joyokus",
			City:   "Malang",
		},
	}

	fmt.Println("Name :",person.Name)
	fmt.Println("Age :",person.Age)
	fmt.Println("Street :",person.Street)
	fmt.Println("City :",person.City)
}

// package main

// import "fmt"

// type Address struct {
// 	Street   string
// 	City     string
// 	Province string
// }

// type PersonEmbbed struct {
// 	Name string
// 	Age  int
// 	Address
// }

// func main() {

// 	address1 := Address{
// 		Street:   "Jatimulyo",
// 		City:     "Bekasi",
// 		Province: "Jawa Barat",
// 	}

// 	person1 := PersonEmbbed{
// 		Name:    "Deni",
// 		Age:     23,
// 		Address: address1,
// 	}

// 	fmt.Println("Name :", person1.Name)
// 	fmt.Println("Age :", person1.Age)
// 	fmt.Println("Street :", person1.Street)
// 	fmt.Println("City :", person1.City)
// }
