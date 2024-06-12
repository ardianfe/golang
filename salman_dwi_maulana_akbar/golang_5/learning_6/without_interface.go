package main

import "fmt"

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof! I am " + d.Name
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow! I am " + c.Name
}

func makeDogSpeak(dog Dog) {
    fmt.Println(dog.Speak())
}

func makeCatSpeak(cat Cat) {
    fmt.Println(cat.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}

    makeDogSpeak(dog)
    makeCatSpeak(cat)
}
