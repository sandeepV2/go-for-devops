package main

import "fmt"

//go:generate stringer -type=Pill
type Pill int
const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol
)



const str = "hello world"
// Untyped constant.
// num can be used to set the value of an int8, int16, int32, int64, uint8, uint16, uint32, or uint64 type
const num = 3
// Typed constant
const num64 int64 = 3

func add(x, y int64) int64 {
    return x + y
}

type specialStr string
func printSpecial(str specialStr){
    fmt.Println(string(str))
}

func main() {
    fmt.Println(add(num, num64)) // Print: 6
	const constHelloWorld = "hello world"
    // var varHelloWorld = "hello world"
    // printSpecial(varHelloWorld) // Won't compile
    printSpecial(constHelloWorld) // Will compile
	fmt.Println(Paracetamol)
}