package main

import (
	"fmt"
	_ "sync"
)

// This led to compile times that were longer than needed and, in some cases, binary sizes that were much bigger than required. Python files were packaged in a proprietary format to ship around production, and some of these unused imports were adding hundreds of megabytes to the files.
// To prevent these types of problems, Go will not compile a program that imports a package but doesn't use it, as shown below:

func main(){
	krishna:= "Hare Krishna!"
	fmt.Println("Here we go!!.")
	fmt.Println(krishna)
}