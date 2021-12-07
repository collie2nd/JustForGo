package main
//
//import "fmt"
//
//func main() {
//	prices := []int{2, 6, 8, 7, 8, 7, 9, 4, 1, 2, 4, 5, 8}
//	k := 3
//	fmt.Println(maxProfit(k, prices))
//}
//
//func maxProfit(k int, prices []int) int {
//	if len(prices) < 2 || k == 0 {
//		return 0
//	}
//	if k >= len(prices) {
//		res := 0
//		for i := 0; i < len(prices)-1; i++ {
//			if prices[i] < prices[i+1] {
//				res += prices[i+1] - prices[i]
//			}
//		}
//		return res
//	}
//	buy := make([]int, k)
//	sell := make([]int, k)
//	buy[0] = -prices[0]
//	sell[0] = 0
//	for i := 1; i < k; i++ {
//		buy[i] = -1 << 62
//		sell[i] = 0
//	}
//	for _, v := range prices {
//		buy[0] = getMax(buy[0], -v)
//		sell[0] = getMax(sell[0], buy[0]+v)
//		for i := 1; i < k; i++ {
//			buy[i] = getMax(buy[i], sell[i-1]-v)
//			sell[i] = getMax(buy[i]+v, sell[i])
//
//		}
//	}
//	return sell[k-1]
//}
//func getMax(a, b int) int {
//	if a < b {
//		return b
//	} else {
//		return a
//	}
//}
