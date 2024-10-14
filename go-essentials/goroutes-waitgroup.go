package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			// decrements the routine
			// exits when reaches o
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	// waits till the task count is 0
	wg.Wait()
	fmt.Println("All work done")
}