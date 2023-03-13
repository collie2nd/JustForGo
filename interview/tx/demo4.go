package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a, b, tmp := 0, 1, 0
	return func() int {
		tmp = a
		a = b
		b = tmp + b
		return a
	}
}
