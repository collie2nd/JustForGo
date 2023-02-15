// https://leetcode.cn/problems/implement-strstr/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("hello", "lo"))
}

// My Code
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 空间复杂度最低算法

// 时间复杂度最低算法
//func strStr(haystack string, needle string) int {
//	i := 0
//	j := 0
//	for ; i < len(haystack); i++ {
//		p := i
//		for j = 0; j < len(needle); j++ {
//			if haystack[p] == needle[j] {
//				p++
//			}
//		}
//		if j == len(needle) {
//			break
//		}
//	}
//	if j == len(needle) {
//		return i
//	} else {
//		return -1
//	}
//}
