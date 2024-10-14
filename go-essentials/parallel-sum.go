package main

// These imports will be used while testing your code

func sumGoRoutine(inputSlice []int, c chan int) {
 //Write your code here
 sum := 0
	for _,v := range inputSlice {
		sum+=v
	}
	c <- sum
}
func findParallelSum(items [][]int) int{
 //Write your code here
 ch := make(chan int)
 rows:= len(items)
 sum:= 0
 for i:=0;i<rows;i++  {
		go sumGoRoutine(items[i], ch)
 }
 
 for j := 0; j < rows; j++ {
		sum += <-ch
	}
 close(ch)	
  
  return sum
 
}