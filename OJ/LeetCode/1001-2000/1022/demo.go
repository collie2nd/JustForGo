// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	fmt.Println()
}

// My Code
//func sumRootToLeaf(root *TreeNode) int {
//	num := 0
//	dfsSumRootToLeaf(root, 0, &num)
//	return num
//}

func dfsSumRootToLeaf(root *TreeNode, a int, b *int) {
	if root == nil {
		return
	}
	a = a*2 + root.Val
	if root.Left == nil && root.Right == nil {
		*b += a
	}
	dfsSumRootToLeaf(root.Left, a, b)
	dfsSumRootToLeaf(root.Right, a, b)
}

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
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return traverse(root, 0)
}

func traverse(root *TreeNode, value int) int {
	sum := 0
	value = (value << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		return value
	}
	if root.Left != nil {
		sum += traverse(root.Left, value)
	}
	if root.Right != nil {
		sum += traverse(root.Right, value)
	}
	return sum
}
