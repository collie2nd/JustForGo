// https://leetcode.cn/problems/length-of-last-word/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello world"))
}

// My Code
//func lengthOfLastWord(s string) int {
//	if len(strings.TrimSpace(s)) == 0 {
//		return 0
//	}
//	var strSlice []string
//	for i := 0; i < len(s); i++ {
//		if s[i] == ' ' {
//			continue
//		} else {
//			if (i > 0 && s[i-1] == ' ') || i == 0 {
//				strSlice = append(strSlice, string(s[i]))
//			} else {
//				strSlice[len(strSlice)-1] += string(s[i])
//			}
//		}
//	}
//	return len(strSlice[len(strSlice)-1])
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
func lengthOfLastWord(s string) int {
	n := len(s)
	right := n - 1
	for ; right >= 0 && s[right] == ' '; right-- {
	}
	left := right
	for ; left >= 0 && s[left] != ' '; left-- {
	}
	return right - left
}
