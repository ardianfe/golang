package main

import (
	"fmt"
)

// Definisi interface sebagai pengganti polymorphism
type Worker interface {
	Work()
}

// Definisi struct Employee sebagai pengganti kelas
// di sini name berawalan huruf kecil yang berarti di enkapsulasi agar tidak bisa diakses seperti ini Employee.name
type Employee struct {
	name string // Private field
	ID   int
	//harus menggunakan getter dan setter untuk menggunakannya
}

// Fungsi NewEmployee untuk membuat instance Employee (constructor)
func NewEmployee(name string, id int) *Employee {
	return &Employee{name: name, ID: id}
}

// Getter untuk name
func (e *Employee) GetName() string {
	return e.name
}

// Setter untuk name
func (e *Employee) SetName(name string) {
	e.name = name
}

// Fungsi Work untuk Employee
func (e Employee) Work() {
	fmt.Println(e.name, "is working")
}

// Definisi struct Manager dengan komposisi sebagai pengganti pewarisan
type Manager struct {
	Employee
}

// Fungsi NewManager untuk membuat instance Manager (constructor)
func NewManager(name string, id int) *Manager {
	return &Manager{Employee: Employee{name: name, ID: id}}
}

// Fungsi Manage untuk Manager
func (m Manager) Manage() {
	fmt.Println(m.GetName(), "is managing")
}

func main() {
	emp := NewEmployee("Alice", 101)
	mngr := NewManager("Bob", 102)

	var w Worker = emp
	w.Work() // Output: Alice is working

	// Polymorphism
	w = mngr
	w.Work()      // Output: Bob is working
	mngr.Manage() // Output: Bob is managing

	// Menggunakan getter dan setter
	fmt.Println("Employee name:", emp.GetName())
	emp.SetName("Alice Smith")
	fmt.Println("Updated employee name:", emp.GetName())
}

// package main

// import "fmt"

// // Definisi interface sebagai pengganti polymorphism
// type Worker interface {
// 	Work()
// }

// // Definisi struct Employee sebagai pengganti kelas
// type Employee struct {
// 	Name string
// 	ID   int
// }

// // Fungsi Work untuk Employee
// func (e Employee) Work() {
// 	fmt.Println(e.Name, "is working")
// }

// // Definisi struct Manager dengan komposisi sebagai pengganti pewarisan
// type Manager struct {
// 	Employee
// }

// // Fungsi Manage untuk Manager
// func (m Manager) Manage() {
// 	fmt.Println(m.Name, "is managing")
// }

// func main() {
// 	emp := Employee{Name: "Alice", ID: 101}
// 	mngr := Manager{Employee: Employee{Name: "Bob", ID: 102}}

// 	var w Worker = emp
// 	w.Work() // Output: Alice is working

// 	// Polymorphism
// 	w = mngr
// 	w.Work()      // Output: Bob is working
// 	mngr.Manage() // Output: Bob is managing
// }
