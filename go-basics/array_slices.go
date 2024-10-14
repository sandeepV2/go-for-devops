package main

import "fmt"



func main(){
	x:= [3]int{}
	y:= changeValueAtZero(x)
	fmt.Println(y)
	// Array are of fixed size.
	// New array for new size, passed to function as new copy.
	s := []int{}
	fmt.Println(len(s))
	someSlice := []int{1, 2, 3}
    for _, val:= range someSlice {
        fmt.Printf("slice entry: %s\n", val)
    }
	for index := range someSlice {
        fmt.Printf("slice entry: %s\n", index)
    }
}

func changeValueAtZero(x [3]int) ([3]int) {
	x[0] = 1
	return x
}