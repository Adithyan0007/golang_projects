package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var number int
	fmt.Scanln(&number)
	fmt.Println(len(number))
}