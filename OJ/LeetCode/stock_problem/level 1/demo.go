package main

import "fmt"

func main() {
	prices := []int{1, 4, 2}
	//prices := []int{7, 6, 4, 3, 1}
	//prices := []int{1, 2, 4}
	fmt.Println(maxProfit(prices))
	fmt.Println(practice(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	gaps := make([]int, len(prices)-1)
	for i := 0; i < len(gaps); i++ {
		gaps[i] = prices[i+1] - prices[i]
	}
	ans := 0
	sum := 0
	for i := 0; i < len(gaps); i++ {
		sum += gaps[i]
		ans = getMax(ans, sum)
		sum = getMax(sum, 0)
	}
	return ans
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func practice(prices []int) int {

	//①要对prices的长度做讨论
	if len(prices) <= 1 {
		return 0
	}
	//②一次make比多次append效率要高
	//gaps := []int{}
	//for i := 0; i < len(prices)-1; i++ {
	//	gaps = append(gaps, prices[i+1]-prices[i])
	//}
	gaps := make([]int, len(prices)-1)
	for i := 0; i < len(gaps); i++ {
		gaps[i] = prices[i+1] - prices[i]
	}

	ans := 0
	max := 0
	//③动态规划，以下错在
	//1. len(gaps)可能等于1,如测试用例[1,2]
	//
	//补充
	//if len(gaps)==1{
	//	if gaps[0]>0{
	//		return gaps[0]
	//	}else {
	//		return 0
	//	}
	//}
	for i := 0; i < len(gaps)-1; i++ {
		max = getMax(max, gaps[i])
		//
		//2. 补充一次更新，不然，会导致第二次更新max，丢失最大的ans，如测试用例[1,4,2]
		//ans = getMax(ans, max)
		//
		max = getMax(gaps[i+1]+max, gaps[i+1])

		ans = getMax(ans, max)
	}
	// ④看上面规范代码，正确理解动态规划
	// sum += gaps[i]
	// ans = getMax(ans, sum)
	// sum = getMax(sum, 0)
	return ans
}
