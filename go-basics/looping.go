// Go differs in that there is a single loop type, for, that can implement the functionality of all the loop types in other languages.
package main

import "fmt"

func main(){
	fmt.Println("Looping")
	for i:=0; i< 5; i++{
		fmt.Println("Basic for loop")
	}
	var j = 10
	for ;j<15; j++{
		fmt.Println("While like looop")
	}
	fmt.Println("Value for j", j)

	b:=true
	var k = 10
	for b {
		fmt.Println("while loop")
		k--
		if k <= 0{
			break
		}
	}
}

