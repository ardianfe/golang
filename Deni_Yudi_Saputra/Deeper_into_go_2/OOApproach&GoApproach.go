package main

import "fmt"

// OOP = Object Oriented Programming

// Dibasic OOP terdapat Encapsulation, Inheritance, Polymorphism, and Abstaction

// Encapsulation = Menyembunyikan detail implementasi dari sebuah class dari akses luar

// Inheritance = Menurunkan sebuah class ke sebuah class lain

// Polymorphism =

// OO Approach Golang

// OOP Ada class kalau golang menggunakan Structs

// Structs

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

//Embedding
// Golang mendukung Embedding yang memungkinkan satu struct untuk menggabungkan field dan metode struct lain, mirip dengan inheritance

type Employee struct {
	Person
	EmployeeID string
}

//Interfaces
//Golang menggunakan interface untuk mendifinisikan kontrak atau perilaku yang harus diimplementasikan oleh tipe yang berbeda.

type Greeter interface {
	Greet() string
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

// Encapsulation di OOP kalau di golang itu pakai struct
// Inheritance di OOP kalau di golang itu pakai embedded yaitu struct dalam struct
// Polymorphism di OOP kalau di golang itu menggunakan interface

// ===== POLYMORPHISM =======

type Vehicle interface {
	Move() string
} // interface yang mendifinisikan satu metode,Move()

type Car struct {
	Brand string
}

func (c Car) Move() string {
	return c.Brand + " car is moving"
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() string {
	return b.Brand + " bicycle is moving"
}

func MoveVehicle(v Vehicle) {
	fmt.Println(v.Move())
}

func main() {
	car :=Car{Brand :"Toyota"}
	bicycle:=Bicycle{Brand:"Polygon"}

	MoveVehicle(car)
	MoveVehicle(bicycle)
}
