package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello"
        ch <- "World"
        close(ch)
    }()

    for msg := range ch {
        fmt.Println(msg)
    }
}

