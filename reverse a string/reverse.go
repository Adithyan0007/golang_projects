package main

import (
	"fmt"
	"strings"
)

func main() {
	reversedstring:= revstring("Adithyan")
	fmt.Print(reversedstring)

}
func revstring(str string) string{
	revs:= strings.Split(str,"")
	data:= ""
	for i:=len(str)-1;i>=0;i--{
		data+=revs[i]
	}
  return data
}