// https://leetcode.cn/problems/search-insert-position/

package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 7}, 3))
}

// My Code
//func searchInsert(nums []int, target int) int {
//	for k, num := range nums {
//		if num >= target {
//			return k
//		}
//	}
//	return len(nums)
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	if target > nums[high] {
		return high + 1
	} else if target == nums[high] {
		return high
	} else if target <= nums[low] {
		return low
	}

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}
