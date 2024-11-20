package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := "this is a string"
	hi()
	hiTow(str)
	n1, n2 := 4, 5
	rs, rm, err := intDivision(n1, n2)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		var st string = strconv.Itoa(rs) + " " + strconv.Itoa(rm)
		hiTow(st)
		fmt.Printf("this is the result : %v, and this is the remainder : %v\n", rs, rm)
	}
	switch {
	case rm == 0:
		fmt.Printf("%v is a divider of %v\n", n2, n1)
	case rs == 0:
		fmt.Printf("%v is equle or grater than %v\n", n2, n1)
	default:
		fmt.Println("everythin simes normal boring")
	}

}

func hi() {
	fmt.Println("go is not that fun is i was expected")
}

func hiTow(par string) {
	fmt.Println("i will print", par)
	fmt.Println(par)
}

func intDivision(n1 int, n2 int) (int, int, error) {
	var err error
	if n2 == 0 {
		err = errors.New("div by zero error")
		return 0, 0, err
	}
	var result int = n1 / n2
	var remainder int = n1 % n2
	return result, remainder, err
}
