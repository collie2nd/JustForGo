// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

package main

import "fmt"

func main() {
	fmt.Println()
}

// My Code
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}

// 空间复杂度最低算法

// 时间复杂度最低算法
