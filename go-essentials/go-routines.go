package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        go fmt.Println(i) // This happens concurrently
    }
    fmt.Println("hello")
    // This is used to prevent the program from exiting
    // before our goroutines above run. We will talk about
    // this later.
    select{}
}