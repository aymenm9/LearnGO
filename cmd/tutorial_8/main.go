package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}
var db = []string{"id1", "id2", "id3", "id4", "id5"}
var results = make([]string, 0)

func main() {
	t0 := time.Now()
	for i := range db {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	wg2.Wait()
	fmt.Printf("total time %v \n", time.Since(t0))
	fmt.Printf("the result %v \n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("val from db", db[i])
	wg2.Add(1)
	go save(db[i])
	wg2.Add(1)
	go log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
	wg2.Done()
}
func log() {
	m.RLock()
	var delay float32 = 1000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("the curnt result is %v \n", results)
	m.RUnlock()
	wg2.Done()
}
func Nlog() {
	m.Lock()
	var delay float32 = 1000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("the curnt result is %v \n", results)
	m.Unlock()
	wg2.Done()
}
