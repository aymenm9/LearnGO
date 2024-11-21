package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5)
	var t0 = time.Now()
	go Pr(c)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		fmt.Println(<-c)
	}
	fmt.Printf("total tiem : %v \n", time.Since(t0))
}

func Pr(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("i finshed")
}
