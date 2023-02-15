// https://leetcode-cn.com/problems/two-sum/

package main

import "fmt"

func main() {
	list1 := []int{9, 9, 1}
	list2 := []int{1}
	var l1, l2 *ListNode
	preNode := &ListNode{}
	l1 = preNode
	for _, v := range list1 {
		oneNode := ListNode{}
		oneNode.Val = v
		oneNode.Next = nil
		preNode.Next = &oneNode
		preNode = preNode.Next
	}

	preNode = &ListNode{}
	l2 = preNode
	for _, v := range list2 {
		oneNode := ListNode{}
		oneNode.Val = v
		oneNode.Next = nil
		preNode.Next = &oneNode
		preNode = preNode.Next
	}
	for p := l1.Next; p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
	fmt.Println()
	for p := l2.Next; p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
	fmt.Println()
	for p := addTwoNumbers(l1.Next, l2.Next); p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
}

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	p := l1
//	q := l2
//	flag := 0
//
//	for {
//		if p != nil && q != nil {
//			sum := p.Val + q.Val + flag
//			if sum >= 10 {
//				p.Val = sum - 10
//				flag = 1
//			} else {
//				p.Val = sum
//				flag = 0
//			}
//		}
//
//		if p.Next == nil && q.Next != nil {
//			p.Next = q.Next
//			for {
//				sum := p.Next.Val + flag
//				if sum >= 10 {
//					p.Next.Val = sum - 10
//					flag = 1
//				} else {
//					p.Next.Val = sum
//					flag = 0
//				}
//				p = p.Next
//				if flag == 0 {
//					break
//				}
//				if p.Next == nil {
//					oneNode := ListNode{
//						1,
//						nil,
//					}
//					p.Next = &oneNode
//					break
//				}
//			}
//			break
//		}
//		if q.Next == nil && p.Next != nil {
//			for {
//				sum := p.Next.Val + flag
//				if sum >= 10 {
//					p.Next.Val = sum - 10
//					flag = 1
//				} else {
//					p.Next.Val = sum
//					flag = 0
//				}
//				p = p.Next
//				if flag == 0 {
//					break
//				}
//				if p.Next == nil {
//					oneNode := ListNode{
//						1,
//						nil,
//					}
//					p.Next = &oneNode
//					break
//				}
//			}
//			break
//		}
//		if q.Next == nil && p.Next == nil {
//			if flag == 1 {
//				oneNode := ListNode{
//					1,
//					nil,
//				}
//				p.Next = &oneNode
//			}
//			break
//		}
//		p = p.Next
//		q = q.Next
//	}
//	return l1
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//链表问题  务必先定义head 空结构体  在定义next指针  next移动即可
	addOne := 0

	var headNode ListNode
	nextNode := &headNode

	for {
		if l1 != nil {
			nextNode.Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			nextNode.Val += l2.Val
			l2 = l2.Next
		}
		nextNode.Val += addOne
		if nextNode.Val > 9 {
			addOne = 1
			nextNode.Val = nextNode.Val - 10
		} else {
			addOne = 0
		}

		if l1 == nil && l2 == nil && addOne == 0 {
			return &headNode
		}
		nextNode.Next = &ListNode{}
		nextNode = nextNode.Next
	}

	return &headNode
}

