// https://leetcode.cn/problems/climbing-stairs/

package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

// My Code
func climbStairs(n int) int {
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = c
		c = b + c
		b = a
	}
	return c
}

// 空间复杂度最低算法

// 时间复杂度最低算法
//func climbStairs(n int) int {
//	sqrt5 := math.Sqrt(5)
//	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
//	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
//	return int(math.Round((pow1 - pow2) / sqrt5))
//}
