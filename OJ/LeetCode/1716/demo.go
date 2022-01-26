// https://leetcode-cn.com/problems/calculate-money-in-leetcode-bank/

package main

import "fmt"

func main() {
	n := 20
	fmt.Println(totalMoney(n))
}

func totalMoney(n int) int {
	a := n/7 + 1
	b := n % 7
	var res float64
	if b == 0 {
		res = 3.5*float64(a-1)*float64(a-1) + 24.5*float64(a-1)
	} else {
		res = 3.5*float64(a)*float64(a) + 17.5*float64(a) + float64(a)*float64(b) - 21 - 0.5*float64(b) + 0.5*float64(b)*float64(b)
	}

	return int(res)
}
