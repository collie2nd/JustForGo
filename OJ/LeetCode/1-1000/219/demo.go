// https://leetcode-cn.com/problems/contains-duplicate-ii/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 1 || len(nums) < 2{
		return false
	}

	numsInfo := make(map[int]int, len(nums))

	for index, v := range nums {
		if preIndex, exist := numsInfo[v]; exist && index-preIndex <= k {
			return true
		}
		numsInfo[v] = index
	}
	return false
}

//空间复杂度最低算法：

//func containsNearbyDuplicate(nums []int, k int) bool {
//	if nums == nil || len(nums) < 2 {
//		return false
//	} else {
//		t := []int{}
//		for i := 0; i < len(nums); i++ {
//			if isContains(t, nums[i]) {
//				return true
//			}
//			t = append(t, []int{nums[i]}...)
//			if len(t) > k {
//				t = append(t[1:])
//			}
//		}
//		return false
//	}
//}
//
//func isContains(nums []int, val int) bool {
//	if len(nums) == 0 {
//		return false
//	} else {
//		for i := 0; i < len(nums); i++ {
//			if nums[i] == val {
//				return true
//			}
//		}
//		return false
//	}
//}

//时间复杂度最低算法：

//func containsNearbyDuplicate(nums []int, k int) bool {
//	if k == 0 || len(nums) < 2 {
//		return false
//	}
//	record := make(map[int]int, len(nums))
//	for i, num := range nums {
//		if j, ok := record[num]; ok && i - j <= k {
//			return true
//		}
//		record[num] = i
//	}
//	return false
//}
