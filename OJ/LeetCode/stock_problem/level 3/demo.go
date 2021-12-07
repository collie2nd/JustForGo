package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 2, 4, 2}
	fmt.Println(maxProfit(prices))
}

const intMin = ^int(^uint(0) >> 1)

func maxProfit(prices []int) int {
	fstBuy, secBuy := intMin, intMin
	fstSell, secSell := 0, 0
	for _, v := range prices {
		fstBuy = getMax(fstBuy, -v)
		fstSell = getMax(fstSell, fstBuy+v)
		secBuy = getMax(secBuy, fstSell-v)
		secSell = getMax(secSell, secBuy+v)
		fmt.Println("fB", fstBuy)
		fmt.Println("fS", fstSell)
		fmt.Println("sB", secBuy)
		fmt.Println("sS", secSell)
		fmt.Println("**************")
	}
	return secSell
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func practice(prices []int) int {
	if len(prices)<=1{
		return 0
	}
	return 0
}

//func maxProfit(prices []int) int {
//	if len(prices) <= 1 {
//		return 0
//	}
//
//	gaps := make([]int, len(prices)-1)
//	mergeGaps := make([]int, 0)
//
//	for i := 0; i < len(gaps); i++ {
//		gaps[i] = prices[i+1] - prices[i]
//	}
//
//	for i := 0; i < len(gaps); i++ {
//		if gaps[i] > 0 {
//			num := gaps[i]
//			for {
//				if i+1 < len(gaps) {
//					if gaps[i+1] > 0 {
//						num += gaps[i+1]
//						i++
//					} else {
//						break
//					}
//				} else {
//					break
//				}
//			}
//			mergeGaps = append(mergeGaps, num)
//		} else {
//			num := gaps[i]
//			for {
//				if i+1 < len(gaps) {
//					if gaps[i+1] <= 0 {
//						num += gaps[i+1]
//						i++
//					} else {
//						break
//					}
//				} else {
//					break
//				}
//			}
//			mergeGaps = append(mergeGaps, num)
//		}
//	}
//
//	//fmt.Println("mergeGaps", mergeGaps)
//
//	profits := make([]int, 0)
//
//	for i := 0; i < len(mergeGaps); i = i + 2 {
//		max1 := getMaxGap(mergeGaps[:i])
//		max2 := getMaxGap(mergeGaps[i:])
//		profits = append(profits, max1+max2)
//	}
//
//	max := getMax(profits)
//
//	return max
//}
//
////所有的正数已经被聚合，所有的非正数也已经被聚合，每隔一个数字就会有一个正数出现
//func getMaxGap(mergeGaps []int) int {
//	num := 0
//	nums := []int{}
//	if len(mergeGaps) > 1 {
//
//		i := 0
//		if mergeGaps[i] <= 0 {
//			i++
//		}
//
//		for ; i < len(mergeGaps); i = i + 2 {
//			for j := 0; i+j < len(mergeGaps); j = j + 2 {
//				for _, v := range mergeGaps[i : i+j+1] {
//					num += v
//				}
//				nums = append(nums, num)
//				num = 0
//			}
//		}
//
//		max := getMax(nums)
//
//		return max
//
//	} else {
//		if len(mergeGaps) == 1 {
//			if mergeGaps[0] > 0 {
//				return mergeGaps[0]
//			} else {
//				return 0
//			}
//		} else {
//			return 0
//		}
//	}
//}
//
//func getMax(nums []int) int {
//	max := 0
//	for _, v := range nums {
//		if max < v {
//			max = v
//		}
//	}
//	return max
//}
