package main

import "fmt"

type Record struct{
    Name string
    Age int
}

func (r Record) String() string {
    return fmt.Sprintf("%s,%d", r.Name, r.Age)
}

func NewRecord(name string, age int) (*Record, error) {
    if name == "" {
        return nil, fmt.Errorf("name cannot be the empty string")
    }
    if age <= 0 {
        return nil, fmt.Errorf("age cannot be <= 0")
    }
    return &Record{Name: name, Age: age}, nil
}

func main() {
    rec, err := NewRecord("John Doak", 100) // Values can be changed here to see the error
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("The name is: ", rec.Name)
    fmt.Println("The age is: ", rec.Age)
}