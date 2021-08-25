package main

import "fmt"

func main() {
	sum := addnums(10, 44)
	fmt.Println(sum)
}

func addnums(x, y int) int {
	return x + y

}