package main

import "fmt"

func main() {
    
	sums([]int{1, 3, 4, 5, 6,7}, findsum)
}

func sums(x []int, a func(int)) {
	
	for _, val := range x {
		a(val)
	}

}
var sum int=0
func findsum(a int) {
	sum = sum + a
	fmt.Println(sum)

}