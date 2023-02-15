// https://leetcode.cn/problems/valid-parentheses/submissions/

package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

// My Code
func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}

	var ret = make([]uint8, 0, length)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			ret = append(ret, s[i])
		case ')':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '(' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		case '}':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '{' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		case ']':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '[' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		}
	}
	return len(ret) == 0
}

// 空间复杂度最低算法

// 时间复杂度最低算法
