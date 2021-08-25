package main

import "fmt"

func main() {
uinp:=enternum()
fmt.Printf("the entered number is %d ",uinp)
}
func enternum() int {
	fmt.Println("enter a number")
	var number int
	fmt.Scanln(&number)
	rem:=0
	sum:=0
	for number!=0 {
	rem = number%10
	sum = sum + rem
	number = number/10

}
 return sum
	
}