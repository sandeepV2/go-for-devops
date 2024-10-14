package main

import (
	"fmt"
	"strings"
	// This import will be used while testing your code
)

func findRepeatingWords(inputString string) string {
 //Write your code here
 inputString = strings.ToLower(inputString)
 wordsinsentence := strings.Fields(inputString)
 counterMap := make(map[string]int)
 repeats := ""
 for idx:= range wordsinsentence {
   word := wordsinsentence[idx]
   word = strings.ReplaceAll(word,",","")
   if _,ok := counterMap[word]; ok{
     counterMap[word] += 1
     if !strings.Contains(repeats, word) {
      repeats += word
     }
   } else {
     counterMap[word] = 1
   }
 }
 fmt.Println(counterMap)
 var wordsList string
 for word,_ := range counterMap{
    if (counterMap[word] > 1) {
      wordsList = wordsList+","+word
    }
  }
  fmt.Println(wordsList)
  if wordsList != "" {
    wordsList = wordsList[1:]
  }
  fmt.Println(wordsList)
 return repeats
}

func main(){
	x := "The rain poured and poured, making the streets flood."
	findRepeatingWords((x))
}