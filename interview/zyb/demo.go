package main

import "fmt"

func main() {
	a := "abcde你好"
	b := []rune(a)
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-1-i] = b[l-1-i], b[i]
	}
	fmt.Println(string(b))
}
