// https://leetcode-cn.com/problems/longest-common-prefix

package main

import (
	"fmt"
)

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

// My Code
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i < len(strs[j]) {
				if strs[0][i] == strs[j][i] {
					continue
				}
				return strs[0][:i]
			} else {
				return strs[j]
			}
		}
	}
	return strs[0]
}

// 空间复杂度最低算法

// 时间复杂度最低算法
