package main

import (
	"fmt"
)

// channel is like pipe between go-routines.
//  Because channels are a pointer-scoped type such as map and slice, we use make()
// send ch <- "word"
// recieve word := <- ch

func main() {
	ch := make(chan string, 1)

	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
		}
		close(ch)
	}()

	for word := range ch {
		fmt.Println(word)
	}
}