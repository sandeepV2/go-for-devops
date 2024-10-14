package main

import "fmt"

type Record struct{
    Name string
    Age int
}

func (r Record) String() string {
    return fmt.Sprintf("%s,%d", r.Name, r.Age)
}

func IncrAge(r *Record) {
    r.Age++
}

func main() {
    // Create a pointer to a Record
    rec := &Record{Age: 15}
    IncrAge(rec)
    fmt.Println("The age is: ", rec.Age)
}