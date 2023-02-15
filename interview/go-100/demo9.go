package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	rand.Seed(time.Now().UnixNano())
	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5)
		}
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
