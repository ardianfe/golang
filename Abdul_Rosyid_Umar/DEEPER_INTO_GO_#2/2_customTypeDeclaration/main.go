package main

import "fmt"

type Name string
type Age int

func main() {
	var myname Name = "abdul"
	var myage Age = 20
	fmt.Printf("my name is %s, \nmy age is %d", myname, myage)
}
