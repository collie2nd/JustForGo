package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Provider(data chan<- int, name string) {
	for i := 0; i < 40; i++ {
		//保证传入i的操作和print操作"原子运行"
		select {
		case data <- i:
			fmt.Println(name, "生产了:", i)
		}
	}
	wg.Done()
}

func Consumer(data <-chan int, name string) {
	for i := 0; i < 40; i++ {
		fmt.Println(name, "读取了", <-data)
	}
	wg.Done()
}

func main() {
	//定义缓冲区大小
	data := make(chan int, 10)
	wg.Add(6)
	go Provider(data, "p1")
	go Provider(data, "p2")
	go Provider(data, "p3")
	go Consumer(data, "c1")
	go Consumer(data, "c2")
	go Consumer(data, "c3")
	wg.Wait()
}
