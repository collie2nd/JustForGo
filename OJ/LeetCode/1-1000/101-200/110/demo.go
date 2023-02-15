// https://leetcode.cn/problems/balanced-binary-tree/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// My Code
//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
//}
//
//func height(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(height(root.Left), height(root.Right)) + 1
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -1 * x
//	}
//	return x
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lHeight := height(root.Left)
		if lHeight == -1 {
			return -1
		}
		rHeight := height(root.Right)
		if rHeight == -1 {
			return -1
		}
		if abs(lHeight-rHeight) > 1 {
			return -1
		}
		return max(lHeight, rHeight) + 1
	}
	return height(root) >= 1
}

func max(nums ...int) int {
	res := math.MinInt
	for k := range nums {
		if nums[k] > res {
			res = nums[k]
		}
	}
	return res
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return num * -1
}
