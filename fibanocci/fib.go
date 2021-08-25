package main

import "fmt"
func printfiban (num int){	
a:=0
b:=1
fmt.Printf("series is : %d %d ",a,b )

for true{
	c:=a+b
	if c >= num{
		break
	}
	fmt.Printf("%d ",c)
	
	
	a=b
	b=c
}}
func main(){
printfiban(10)

}