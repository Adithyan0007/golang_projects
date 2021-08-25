package main

import "fmt"

func main() {
	num := 17

	for i := num; i > 0; i-- {
		if i == 1 || i == num {
			continue
		}
		if num % i == 0 {
			fmt.Printf("not a prime")
		}else if i==2{
			fmt.Println("prime number")
		}
	}
	
}