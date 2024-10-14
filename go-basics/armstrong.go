package main

import (
	"fmt"
	"math"
)

// These imports will be used while testing your code

func isArmstrong(num int) string {
    // Complete the function signature and write your code here
    if num > 999 || num < 100 { 
      return "InputError"
    }
    var res float64
    iter:= num
    for iter > 0 {
      rem:= iter % 10
      iter = iter / 10
      res += math.Pow(float64(rem), 3)
    }
	fmt.Println(res);
    if res == float64(num) {
      return "Armstrong"
    } else {
      return "NotArmstrong"
    }
    
}

func main(){
	fmt.Println(isArmstrong(371))
}