// Package scoped: Can be seen by the entire package and is declared outside a function

// Function scoped: Can be seen within {} which defines the function

// Statement scoped: Can be seen within {} of a statement in a function (for loop, if/else)

package main

import "fmt"

var pckg_scoped = "hello"

func main() {
    func_scoped := "function scoped"
    fmt.Println(pckg_scoped)
    fmt.Println(func_scoped)

    // statement scoped
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}

// The rule here is that if we create a variable within a function or statement, it must be used. This is much for the same reason as package imports; declaring a variable that isn't used is almost always a mistake.

// This can be relaxed in much the same way as an import, using _, but is far less common.

// needed, _ := someFunc() 

// Most commonly used when a function returns multiple values, but we only need one. Here, we create and assign to the needed variable, but the second value isn't something we use, so we drop it.