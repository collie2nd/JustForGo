package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{11: 1}
	fmt.Println(input, len(input))
	handle(input)
}

func handle(input []int) {
	var wg sync.WaitGroup
	apiCh := make(chan struct{}, 10)
	wg.Add(len(input))
	for _, v := range input {
		apiCh <- struct{}{}
		go func(v int) {
			fmt.Println(v) // handle function
			time.Sleep(time.Second)
			wg.Done()
			<-apiCh
		}(v)
	}
	wg.Wait()
	return
}
