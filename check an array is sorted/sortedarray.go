package main

import "fmt"
func checksortedarray(arr []int){
   sortedArray := true
   for i:=0; i<=len(arr)-1; i++{
      for j:=0; j<len(arr)-1-i; j++{
         if arr[j]> arr[j+1]{
            sortedArray = false
            break
         }
      }
   }
   if sortedArray{
      fmt.Println("Given array is already sorted.")
   } else {
      fmt.Println("Given array is not sorted.")
   }
}

func main(){
   checksortedarray([]int{1, 3, 5, 6, 7, 8})
}