// https://leetcode.cn/problems/defanging-an-ip-address/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("192.168.1.1"))
}

// My Code
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

// 空间复杂度最低算法

// 时间复杂度最低算法
