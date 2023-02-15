// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// My Code
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 空间复杂度最低算法

// 时间复杂度最低算法
