// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 4, 4, 5, 6}))
}

// My Code
func removeDuplicates(nums []int) int {
	slow := 1
	l := len(nums)
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 空间复杂度最低算法

// 时间复杂度最低算法
