package main

import "fmt"

func main() {
    var i interface{}
    i = "hello world"
    if v, ok := i.(string); ok {
        fmt.Println(v)
    }
}