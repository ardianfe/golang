package main

import "fmt"

// ====FUNCTION BIASA =======
func sayHello() {
	fmt.Println("Takitaki")
}

// ======FUNCTION PARAMETER=======
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// =====FUNCTION RETURN VALUE=======
// function mengembalikan data

func getHello(name string) string {

	hello := "Hello" + name
	return hello
}

func add(a int, b int) int {
	return a + b
}

// =====FUNCTION MULTIPLE RETURN VALUE=======

func getName() (string, string) {
	return "deni", "yudiku"
}

// (int,int) adalah hasil returnnya nanti, (a int, b int) adalah parameternya
func divide(a int, b int) (int, int) {
	tambah := a + b
	bagi := a / b

	return tambah, bagi

}

//  dengan menamai value pada returnan, parameter bisa disingkat dengan seperti dibawah dengan int di belakang sama saja dengan (quotient int, remainder int)
func divide1(a int, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a - b
	return // Nama return values digunakan secara implisit
}

// anonim return dengan int pada return

func anonim(a int, b int) int {
	return a - b
}

//anonim return parameter
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// ======= MAIN FUNCTION =========

func main() {
	sayHello()
	sayHelloTo("Deni", "Yudi")
	// return value
	fmt.Println(getHello("Saputra"))
	sum := add(1, 3)
	fmt.Println(sum)

	// multiple return value,

	firstname, lastname := getName() //misal gak diambil salah satu atau iqnore maka pakai "_"
	fmt.Println(firstname, lastname)

	t, b := divide(10, 2)
	fmt.Println("tambah", t, "bagi", b)

	q, r := divide1(10, 5)
	fmt.Println("bagi", q, "kurang", r)

	// anonymouse return dengan tidak menamai return

	min := anonim(3, 2)
	fmt.Println("anon kurang = ", min)

	// anonymouse return parameter
	add := func(a int, b int) int {
		return a + b
	}

	result := applyOperation(3, 4, add)
	fmt.Println("Result:", result)

}
