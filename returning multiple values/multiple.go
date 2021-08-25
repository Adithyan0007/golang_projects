package main

import "fmt"

func main() {
	sum, product := nums(10, 34)

	fmt.Printf("sum is %v and product is %v",sum,product)
}
func nums(num1, num2 int) (int, int) {
	sum := num1 + num2
	product := num1 * num2

	return sum, product

}

