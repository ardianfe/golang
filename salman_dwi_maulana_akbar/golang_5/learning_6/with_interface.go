package main

import "fmt"

// Speaker interface defines a method Speak
type Speaker interface {
    Speak() string
}

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

type Bird struct {
    Name string
}

func (b Bird) Speak() string {
    return "Tweet! I am " + b.Name
}

// makeSpeak function takes a Speaker interface
func makeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    bird := Bird{Name: "Tweety"}

    makeSpeak(dog)
    makeSpeak(cat)
    makeSpeak(bird)
}
