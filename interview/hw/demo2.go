package main

import (
	"fmt"
)

func main() {
	var m, n int
	var fields []int
	fmt.Scan(&m, &n)
	for i := 0; i < m; i++ {
		var f int
		fmt.Scan(&f)
		fields = append(fields, f)
	}

	if n < m {
		fmt.Println(-1)
	} else if n == m {
		fmt.Println(getMaxInt(fields))
	} else {
		for k := 1; k <= getMaxInt(fields); k++ {
			var sumDays int
			for _, f := range fields {
				if f%k == 0 {
					sumDays += f / k
				} else {
					sumDays += f/k + 1
				}
			}
			if sumDays <= n {
				fmt.Println(k)
				return
			}
		}
	}
}

func getMaxInt(list []int) int {
	var max int
	for _, v := range list {
		if max < v {
			max = v
		}
	}
	return max
}
