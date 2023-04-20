package main

import "fmt"

func main() {
	f := fbHandler()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

func fbHandler() func() int {
	a, b, tmp := 0, 1, 0
	return func() int {
		tmp = a
		a = b
		b = tmp + b
		return a
	}
}
