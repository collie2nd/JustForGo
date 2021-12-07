package main

import (
	"fmt"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change(amount,coins))
}

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]]
		}
	}
	return dp[amount]
}
