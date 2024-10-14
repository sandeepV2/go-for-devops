package main

import (
	"fmt"
)

func reverse_int(n int) int {
    new_int := 0
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder 
        n /= 10
    }
    return new_int 
}

func main() {
    fmt.Println(reverse_int(123456))
    fmt.Println(reverse_int(100))
    fmt.Println(reverse_int(1001))
    fmt.Println(reverse_int(131415))
    fmt.Println(reverse_int(1357))
}