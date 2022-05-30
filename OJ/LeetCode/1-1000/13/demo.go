// https://leetcode-cn.com/problems/roman-to-integer/

package main

import (
	"fmt"
)

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

// My Code
//func romanToInt(s string) int {
//	sum := 0
//	for i := 0; i < len(s); i++ {
//		switch s[i] {
//		case 'M':
//			sum += 1000
//		case 'D':
//			sum += 500
//		case 'C':
//			if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
//				sum -= 100
//				break
//			}
//			sum += 100
//		case 'L':
//			sum += 50
//		case 'X':
//			if i+1 < len(s) && (s[i+1] == 'C' || s[i+1] == 'L') {
//				sum -= 10
//				break
//			}
//			sum += 10
//		case 'V':
//			sum += 5
//		case 'I':
//			if i+1 < len(s) && (s[i+1] == 'X' || s[i+1] == 'V') {
//				sum -= 1
//				break
//			}
//			sum += 1
//		}
//	}
//	return sum
//}

// 空间复杂度最低算法

// 时间复杂度最低算法
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
