// https://leetcode.cn/problems/sqrtx/

package main

import "fmt"

func main() {
	fmt.Println()
}

// My Code
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 空间复杂度最低算法

// 时间复杂度最低算法
