package main

import "fmt"

func main() {
	num := 6
	fact := 1
	for num != 0 {
		fact = fact * num
		num--
	}
	fmt.Println(fact)
}