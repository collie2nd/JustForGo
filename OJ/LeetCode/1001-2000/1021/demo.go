// https://leetcode.cn/problems/remove-outermost-parentheses/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}

// My Code
//func removeOuterParentheses(s string) string {
//	var rst, stack string
//	var head int
//	for i, e := range s {
//		if e == '(' {
//			if len(stack) == 0 {
//				head = i + 1
//				stack += "("
//			} else if len(stack) != 0 && stack[len(stack)-1] == '(' && e == '(' {
//				stack += "("
//			}
//		} else if e == ')' {
//			if len(stack) > 1 && stack[len(stack)-1] == '(' && e == ')' {
//				stack = stack[:len(stack)-1]
//			} else if len(stack) == 1 && stack[len(stack)-1] == '(' && e == ')' {
//				stack = stack[:len(stack)-1]
//				rst += s[head:i]
//			}
//		}
//	}
//	return rst
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
func removeOuterParentheses(s string) string {
	b := 0
	r := make([]byte, 0, len(s)-2)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			b++
			if b > 1 {
				r = append(r, '(')
			}
		} else {
			b--
			if b > 0 {
				r = append(r, ')')
			}
		}
	}
	return string(r)
}
