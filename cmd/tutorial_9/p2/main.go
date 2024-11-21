package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXCHICKENPRICE float32 = 5.0
const GOODDEAL float32 = 4.5

func main() {
	var chickenchannel = make(chan string)
	var goodDealchannel = make(chan string)
	var websites = []string{"walmart.com", "castco.net", "wholefoods.com"}
	t0 := time.Now()
	for _, w := range websites {
		go checkChickenPrices(w, chickenchannel)
		go checkgoodDeal(w, goodDealchannel)
	}
	sendMessage(chickenchannel, goodDealchannel)
	fmt.Printf("total time : %v\n", time.Since(t0))
}

func checkChickenPrices(website string, chickenchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 100
		if chickenPrice <= MAXCHICKENPRICE {
			chickenchannel <- website
			break
		}
	}
}
func checkgoodDeal(website string, goodDealchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var Price = rand.Float32() * 20
		if Price <= GOODDEAL {
			goodDealchannel <- website
			break
		}
	}
}

func sendMessage(chickenchannel chan string, goodDealchannel chan string) {
	select {
	case <-chickenchannel:
		fmt.Printf("Found a deal on chicken at %s\n", <-chickenchannel)
	case <-goodDealchannel:
		fmt.Printf("Found a deal on anthing at %s\n", <-goodDealchannel)
	}

}
