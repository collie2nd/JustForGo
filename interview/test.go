package main

import "fmt"

func main() {
	str := "abcd"
	getSubStr([]byte(str))
}
func getSubStr(str []byte) {
	if len(str) == 0 {
		return
	}

	if len(str) > 0 {
		i := 0
		l := 1
		max := len(str)
		for l <= max {
			if i+l-1 < max {
				fmt.Println(string(str[i : i+l]))
				i++
			} else {
				i = 0
				l++
			}
		}
	}
	return
}
