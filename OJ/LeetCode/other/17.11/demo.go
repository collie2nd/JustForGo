// https://leetcode.cn/problems/find-closest-lcci/

package main

import (
	"fmt"
	"math"
)

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	fmt.Println(findClosest(words, "student", "a"))
}

// My Code
func findClosest(words []string, word1 string, word2 string) int {
	p := -1
	q := -1
	max := len(words)
	for index, word := range words {
		if word == word1 {
			p = index
		} else if word == word2 {
			q = index
		}
		tmp := len(words)
		if p != -1 && q != -1 {
			tmp = int(math.Abs(float64(q - p)))
		}
		if tmp < max {
			max = tmp
		}
	}
	return max
}

// 空间复杂度最低算法
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//
//func findClosest(words []string, word1 string, word2 string) int {
//	ans := int(1e9)
//	p1, p2 := -1, -1
//	for i, w := range words {
//		if w == word1 {
//			p1 = i
//		} else if w == word2 {
//			p2 = i
//		}
//		if p1!=-1 && p2!=-1 {
//			ans = min(ans, abs(p1-p2))
//		}
//	}
//	return ans
//}

// 时间复杂度最低算法
//func findClosest(words []string, word1 string, word2 string) int {
//	result, x, y := len(words)-1, -1, -1
//	abs := func(num int) int {
//		if num < 0 {
//			return -num
//		}
//		return num
//	}
//	for i := range words {
//		if words[i] == word1 {
//			x = i
//		}
//		if words[i] == word2 {
//			y = i
//		}
//		if x != -1 && y != -1 && abs(x-y) < result {
//			result = abs(x - y)
//			if result == 1 {
//				return 1
//			}
//		}
//	}
//	return result
//}
