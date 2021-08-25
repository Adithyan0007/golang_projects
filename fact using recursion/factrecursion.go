package main

import "fmt"

func main() {
	nums := factor(5)
	fmt.Println(nums)
}

func factor(num int) int {
	if num==0 || num==1{
		return num
	}
	return num * factor(num-1)
}