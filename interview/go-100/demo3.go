package main

import (
	"fmt"
)

func main() {
	s := "adsfger4刘星"
	fmt.Println(revertStr(s))
}

func revertStr(str string) string {
	s := []rune(str)
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return string(s)
}
