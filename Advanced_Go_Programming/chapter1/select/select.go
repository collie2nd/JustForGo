package main

import (
	"fmt"
	"time"
)

//func main() {
//	ch := make(chan int)
//	go func() {
//		for {
//			select {
//			case ch <- 0:
//			case ch <- 1:
//			}
//		}
//	}()
//	for v := range ch {
//		fmt.Println(v)
//	}
//}

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-cannel:
		}
	}
}

func main() {
	start := time.Now()
	cannel := make(chan bool)
	go worker(cannel)

	time.Sleep(time.Second)
	cannel <- true
	fmt.Println(time.Since(start))
}
