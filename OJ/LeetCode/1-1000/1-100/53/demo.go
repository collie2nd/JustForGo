// https://leetcode.cn/problems/maximum-subarray/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// My Code
//func maxSubArray(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	for k, num := range nums {
//		if k == 0 {
//			continue
//		}
//		if nums[k-1] > 0 {
//			nums[k] = nums[k-1] + num
//		}
//	}
//
//	//fmt.Println(nums)
//	max := nums[0]
//	for _, num := range nums {
//		if num > max {
//			max = num
//		}
//	}
//	return max
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
func maxSubArray(nums []int) int {
	// 动态规划
	result := math.MinInt32
	numsSize := len(nums)
	dp := nums[0]
	result = dp
	for i := 1; i < numsSize; i++ {
		// 当前值和前一次最大值相加，然后和当前值做比较，取最大值
		dp = maxInt(dp+nums[i], nums[i])
		// 上一步得到的最大值再与上一个循环得到的最大值比较，取最大值
		result = maxInt(result, dp)
	}

	return result
	/* 贪心法
	    	result := math.MinInt32
		numsSize := len(nums)
		sum := 0

		for i:=0; i<numsSize; i++ {
			// 累加当前的元素
			sum += nums[i]
			// 得到累加当前元素后的最大值，并通过maxInt函数得到最大值
			result = maxInt(result, sum);
			//如果sum < 0，重新开始找子序串
			//(你们前面干了半天还是负债累累，我退出，从头开始，及时止损，开启新的人生)
			if sum < 0 {
				sum = 0
			}
		}

		return result
	*/
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
