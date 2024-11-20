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
	var head *node = nil
	for i := 0; i < len(arrPtr); i++ {
		head = appendElment(arrPtr[i]+1, head)
	}
	printList(head)
}
func creatElement(val int8) node {
	return node{val, nil}
}
func appendElment(v int8, h *node) *node {
	n := creatElement(v)
	n.next = h
	return &n
}

func printList(h *node) {
	if h == nil {
		return
	}
	printList(h.next)
	fmt.Println(h.val)
}
