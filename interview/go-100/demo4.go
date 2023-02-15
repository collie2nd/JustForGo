package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "adsfger4"
	s2 := "asfgder4"
	fmt.Println(isRegroup(s1, s2))
}

func isRegroup(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, v := range str1 {
		if strings.Count(str1, string(v)) != strings.Count(str2, string(v)) {
			return false
		}
	}
	return true
}
