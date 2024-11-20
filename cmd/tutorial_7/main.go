package main

import "fmt"

type node struct {
	val  int8
	next *node
}

func main() {
	var arrPtr *[10]int8 = new([10]int8)
	var ptr *int8
	for i := 0; i < 10; i++ {
		ptr = &arrPtr[i]
		*ptr = int8(i)
	}
	fmt.Println(arrPtr)
}

/*func creatList(val int8) node {
	return node{val, nil}
}*/
