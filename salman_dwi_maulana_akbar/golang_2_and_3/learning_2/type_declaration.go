package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
    var tempC Celsius = 36.5
    var tempF Fahrenheit = 97.7

    fmt.Println("Suhu dalam Celsius:", tempC)
    fmt.Println("Suhu dalam Fahrenheit:", tempF)
}
