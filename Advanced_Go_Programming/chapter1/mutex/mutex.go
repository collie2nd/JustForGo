package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

var total2 uint64

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go worker(&wg)
	go worker(&wg)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()
	fmt.Println(total.value)
	fmt.Println(total2)
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}
