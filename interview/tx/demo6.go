package main

import (
	"fmt"
)

func main() {
	fmt.Println(judge([]int{13, 12, 11, 0, 1}))
	fmt.Println(judge([]int{6, 0, 2, 0, 4}))
}

//
//func judge(nums []int) bool {
//	sort.Ints(nums)
//	max := nums[4]
//	var king int
//	for i := 0; i < 4; i++ {
//		if nums[i] == 0 {
//			king++
//		} else if nums[i] == nums[i+1] {
//			return false
//		}
//	}
//	return max-nums[king] <= 5
//}

func judge(nums []int) bool {
	var king int
	max := -1
	min := 14
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min && v != 0 {
			min = v
		}
		if v == 0 {
			king++
		}
	}
	fmt.Println(max, min, king, 5-king)
	return max-min+1 > 5-king
}
