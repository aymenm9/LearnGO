package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int8
	intNum++
	fmt.Println(intNum)
	var floatVar float32 = 50.22
	// intNum + floatVar : error
	var r float64 = float64(floatVar) + float64(intNum)
	fmt.Println(r)

	var str string = "js is bad" + " go instaltion is horeble " + "python is the best"
	fmt.Println(str)
	fmt.Println(utf8.RuneCountInString(str))
	var char rune = 'a'
	fmt.Println("char val is :", char)

	newVar := "my var" // var newVar string = "my var"

	fmt.Println(newVar)
}
