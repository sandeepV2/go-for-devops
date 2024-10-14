// Maps are like dictionaries in python.

package main

import "fmt"

func main() {
	var counters = make(map[string]int, 10) 

	fmt.Println(counters, len(counters));
	counters["ab"] = 1
	for k,v := range counters {
		fmt.Println(k,v)
	}

	modelToMake := map[string]string{
        "prius": "toyota",
        "chevelle": "chevy",
    }

    if carMake, ok := modelToMake["outback"]; ok {
        fmt.Printf("car model \"outback\" has make %q", carMake)
    }else{
        fmt.Printf("car model \"outback\" has an unknown make")
    }
}