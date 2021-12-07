package main

import "fmt"

func main() {
	prices := []int{2, 6, 8, 7, 8, 7, 9, 4, 1, 2, 4, 5, 8}
	k := 3
	fmt.Println(maxProfit(k, prices))
}

func profits(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func maxProfit(k int, prices []int) int {
	if prices == nil || len(prices) == 0 || k <= 0 {
		return 0
	}
	if k >= len(prices) {
		return profits(prices)
	}

	ownDp := make([][]int, 2)
	noDp := make([][]int, 2)
	for i := range ownDp {
		ownDp[i] = make([]int, k+1)
		noDp[i] = make([]int, k+1)
	}

	for t := 0; t <= k; t++ {
		noDp[0][t] = 0
		ownDp[0][t] = -prices[0]
	}
	ownDp[0][k] = -1<<32 - 1
	res := 0
	for i := 1; i < len(prices); i++ {
		x := i % 2
		y := (i - 1) % 2
		for m := 0; m <= k; m++ {
			if m == 0 {
				noDp[x][m] = max(-1<<32-1, noDp[y][m])
			} else {
				noDp[x][m] = max(noDp[y][m], ownDp[y][m-1]+prices[i])
			}
			ownDp[x][m] = max(ownDp[y][m], noDp[y][m]-prices[i])
			tmp := max(noDp[x][m], ownDp[x][m])
			res = max(res, tmp)
		}

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
