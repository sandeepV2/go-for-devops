package main

import "fmt"

func sum(numbers []int) int {
    sum := 0
    for _, n := range numbers {
        sum += n  
    }
    return sum
}

func main() {
    args := []int{1,2,3,4,5}
    fmt.Println(sum(args))
	fmt.Println(sum1(1,2,3,4,5));
	// Use of anonymous functions.
	res,_ := func(word1, word2 string) (w1,w2 string) {
        return word1, word2
    }("hello", "world")
    fmt.Println(res) 
}

// variadic notation or spread operator in javascript.
func sum1(numbers ...int) int {
    sum := 0
    for _, n := range numbers {
        sum += n  
    }
    return sum
}
