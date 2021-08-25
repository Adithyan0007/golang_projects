package main

import (
	"fmt"
	"math"
)

func main() {
area(10.233)

}
func area(radius float64){
	res:= math.Pi*radius*radius
	fmt.Println(res) 
	
}