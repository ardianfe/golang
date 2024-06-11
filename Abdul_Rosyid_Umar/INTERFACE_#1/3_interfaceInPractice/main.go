package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	HitungLuas() float64
	HitungKeliling() float64
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) HitungLuas() float64 {
	return 3.14 * l.JariJari * l.JariJari
}
func (l Lingkaran) HitungKeliling() float64 {
	return 2 * 3.14 * l.JariJari
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

type Segitiga struct {
	Alas, Tinggi, SisiA, SisiB, SisiC float64
}

func (p PersegiPanjang) HitungLuas() float64 {
	return p.Panjang * p.Lebar
}

func (s Segitiga) HitungLuas() float64 {
	return s.Alas * s.Tinggi * 0.5
}

func (p PersegiPanjang) HitungKeliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}

func (s Segitiga) HitungKeliling() float64 {
	return s.SisiA + s.SisiB + s.SisiC
}

func PrintDetails(shape Shape) {
	t := reflect.TypeOf(shape).Name()
	fmt.Printf("Detail Bangun datar: %s\n", t)
	fmt.Printf("Luas: %v\n", shape.HitungLuas())
	fmt.Printf("Keliling: %v\n", shape.HitungKeliling())
}

func main() {
	circle := Lingkaran{JariJari: 7}
	PrintDetails(circle)

	rectangle := PersegiPanjang{Panjang: 10, Lebar: 5}
	PrintDetails(rectangle)

	tringle := Segitiga{Alas: 6, Tinggi: 8, SisiA: 6, SisiB: 8, SisiC: 10}
	PrintDetails(tringle)
}
