//https://leetcode.cn/problems/add-binary/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(addBinary("1", "11"))
}

// My Code
func addBinary(a string, b string) string {

	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)

	ai.Add(ai, bi)
	return ai.Text(2)
}

// 空间复杂度最低算法

// 时间复杂度最低算法
