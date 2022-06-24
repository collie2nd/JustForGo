// https://leetcode.cn/problems/remove-element/

package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{1, 1, 2, 3, 4, 4, 4, 5, 6}, 4))
}

// My Code
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 空间复杂度最低算法

// 时间复杂度最低算法
