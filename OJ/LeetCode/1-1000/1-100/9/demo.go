// https://leetcode-cn.com/problems/palindrome-number/

package main

import (
	"fmt"
)

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}

// My Code
//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	xStr := strconv.Itoa(x)
//	i := 0
//	j := len(xStr) - 1
//	for {
//		if i < j {
//			if xStr[i] == xStr[j] {
//				i++
//				j--
//			} else {
//				return false
//			}
//		} else {
//			break
//		}
//	}
//	return true
//}

// 空间复杂度最低算法
func isPalindrome(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}

	return x == y || x == y/10
}

// 时间复杂度最低算法
