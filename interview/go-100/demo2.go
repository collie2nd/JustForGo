package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "adsfger4"
	fmt.Println(strings.Count(s, "") - 1)
	fmt.Println(isUniqueStr(s))
}

func isUniqueStr(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}
