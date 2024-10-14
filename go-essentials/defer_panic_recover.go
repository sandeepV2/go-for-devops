package main

import (
	"fmt"
	"log"
)

//defer is executed after pgm. cleanup.
//panic at exit
//recover for panic in gRPC.

func someFunc() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("called recover, panic was: %q", r)
        }
    }()
    panic("oh no!!!")
}


func printStuff() (value string) {
    defer fmt.Println("exiting")
    defer func() {
        value = "we returned this"
    }()
    fmt.Println("I am printing stuff")
    return ""
}


func main() {
	v:= printStuff()
	fmt.Println(v)
	someFunc()
}