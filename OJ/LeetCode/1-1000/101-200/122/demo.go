// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(practice(prices))
}

func maxProfit(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {

		if prices[i+1] > prices[i] {
			ans += prices[i+1] - prices[i]
		}
	}
	return ans
}

func practice(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		//可以优化存储空间,省去gap的分配空间，直接对比   price[i+1]>price[i]
		//然后   ans += price[i+1] - price[i]
		gap := prices[i+1] - prices[i]
		if gap > 0 {
			ans += gap
		}
	}
	return ans
}
