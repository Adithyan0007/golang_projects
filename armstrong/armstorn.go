package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var num int
	fmt.Scanln(&num)
	var arm int
	tmp:=num
	for num!=0{
		rem:=num%10
		arm = arm + (rem*rem*rem)
		num = num/10
	}
	if tmp == arm{
		fmt.Println("armstrong")
	}else{
		fmt.Println("not armstrong")
	}
}