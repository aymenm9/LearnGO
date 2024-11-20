package main

import "fmt"

type car interface {
	ratecar()
}

type ev struct {
	name    string
	carType string
}
type realCar struct {
	name    string
	carType string
}

func (c realCar) ratecar() {
	fmt.Println(c.name, "you are good to go")
}

func (c ev) ratecar() {
	fmt.Println(c.name, "bay a real car")
}
func main() {
	var c1 car = realCar{"g63", "realcar"}
	var c2 car = ev{"modl x", "ev"}
	c1.ratecar()
	c2.ratecar()
}
