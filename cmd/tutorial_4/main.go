package main

import "fmt"

func main() {

	var intArr [3]int8
	intArr[0] = 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])
	fmt.Println(&intArr[0], &intArr[1])
	fmt.Println(len(intArr))
	arr := [...]int8{1, 5}
	fmt.Println(arr)
	var sliceArr []int8 = []int8{1, 2, 3}
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr), &sliceArr[0])
	sliceArr = append(sliceArr, 4)
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr), &sliceArr[0])
	var sarr []int8 = make([]int8, 5, 6)
	sarr = append(sarr, 6)
	fmt.Println(sarr, len(sarr), cap(sarr), &sarr[0])
	sarr = append(sarr, 7)
	fmt.Println(sarr, len(sarr), cap(sarr), &sarr[0])
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var mapTow map[uint8]string = map[uint8]string{1: "Aymen"}
	fmt.Println(mapTow)
	myMap["aymen"] = 1
	fmt.Println(myMap)
	fmt.Println(myMap["Aymen"])
	fmt.Println(myMap)
	delete(myMap, "aymen")
	fmt.Println(myMap)
	for i, v := range sliceArr {
		fmt.Println(i, v)
	}

}
