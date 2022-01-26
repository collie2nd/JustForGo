package _001

import (
	"fmt"
)

func Callatz()  {
	var n, path int
	fmt.Scanf("%v", &n)
	if n <= 1000 && n > 1 {
		for {
			if n%2 == 1 {
				n = (3*n + 1) / 2
			} else {
				n = n / 2
			}
			path++
			if n == 1 {
				break
			}
		}
	}
	fmt.Printf("%v", path)
	return
}
