package main

import "fmt"

func someFunc() {
    defer func() {
        recover()
    }()
    panic("Error!")
}

func main() {
    someFunc()
    fmt.Print("The End!")
}