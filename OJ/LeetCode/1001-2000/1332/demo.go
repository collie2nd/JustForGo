// https://leetcode-cn.com/problems/remove-palindromic-subsequences/

package main

import (
	"fmt"
)

func main() {
	stones := "abbaaaab"
	fmt.Println(removePalindromeSub(stones))
}

//

func removePalindromeSub(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	if isPalindrome(s) {
		return 1
	} else {
		return 2
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); {
		for j := len(s) - 1; j >= 0; {
			if i < j && s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
			if i >= j {
				return true
			}
		}
	}
	return true
}
