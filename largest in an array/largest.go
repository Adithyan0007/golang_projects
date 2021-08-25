package main

import "fmt"
func large(num int) int {
	arr := []int{1, 3, 3, 2, 12432, 54, 32, 4}
	arr = append(arr, num)
	larges := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > larges {
			larges = arr[i]
		}

	}
	return larges
}
func main() {
	nums := large(1999)
	fmt.Println(nums)
}
