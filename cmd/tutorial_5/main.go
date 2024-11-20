package main

import (
	"fmt"
	"strings"
)

func main() {
	var myStr string = "résumé"
	var char = myStr[1:3]
	fmt.Printf("%s, %T\n", char, char)
	for i, v := range myStr {
		fmt.Printf("index: %v, loc: %v val: %v rval: %c \n", i, &v, v, v)
	}
	var strs []string = []string{"aymen", "is", "learning", "go"}
	var strBuilder strings.Builder
	for i := range strs {
		strBuilder.WriteString(strs[i])
		strBuilder.WriteString(" ")
	}
	fmt.Println(strBuilder.String())
}
