package main

import "fmt"


func Greet(name string) (string, error){
	if name == "" {
		return "", fmt.Errorf("name was empty!")
	}

	return "Hello "+name,nil
}

func main(){
	fmt.Println(Greet("Sandeep"))
}