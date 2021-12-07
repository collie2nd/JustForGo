package main

import (
	"fmt"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for v := 1; v <= amount; v++ {
		dp[v] = 0x3f3f3f3f
	}
	for _, coin := range coins {
		for v := coin; v <= amount; v++ {
			if dp[v] > dp[v-coin] + 1 {
				dp[v] = dp[v-coin] + 1
			}
		}
	}
	if dp[amount] == 0x3f3f3f3f {
		return -1
	}
	return dp[amount]
}