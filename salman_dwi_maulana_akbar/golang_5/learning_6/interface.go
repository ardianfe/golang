package main

import "fmt"

type Transport interface {
	GetDetails() string
	Book() string
}

type Car struct {
	Model string
	Seats int
}

func (c Car) GetDetails() string {
	return fmt.Sprintf("Car Model: %s, Seats: %d", c.Model, c.Seats)
}

func (c Car) Book() string {
	return fmt.Sprintf("Car %s booked!", c.Model)
}

type Bus struct {
	Route string
	Capacity int
}

func (b Bus) GetDetails() string {
	return fmt.Sprintf("Bus Route: %s, Capacity: %d", b.Route, b.Capacity)
}

func (b Bus) Book() string {
	return fmt.Sprintf("Bus on route %s booked!", b.Route)
}

type Bike struct {
	Brand string
	Type  string
}

func (bi Bike) GetDetails() string {
	return fmt.Sprintf("Bike Brand: %s, Type: %s", bi.Brand, bi.Type)
}

func (bi Bike) Book() string {
	return fmt.Sprintf("Bike %s (%s) booked!", bi.Brand, bi.Type)
}

func PrintDetails(t Transport) {
	fmt.Println(t.GetDetails())
}

func BookTransport(t Transport) {
	fmt.Println(t.Book())
}

func main() {
	car := Car{Model: "Toyota Corolla", Seats: 4}
	bus := Bus{Route: "Downtown", Capacity: 50}
	bike := Bike{Brand: "Trek", Type: "Mountain"}

	// Menggunakan fungsi dengan parameter bertipe Transport
	PrintDetails(car)
	PrintDetails(bus)
	PrintDetails(bike)

	BookTransport(car)
	BookTransport(bus)
	BookTransport(bike)
}
