//对于一个字节（8bit）的无符号整型变量，求其二进制表示中“1”的个数。
//比如：10000001；1的个数位2
//
//            11000001；1的个数位3
//
//求f(N)
//
//说出思路并写出该函数，给出对应的时间复杂度，怎么保证效率最高？
//
//python和golang中使用0b表示二进制，例如 0b10000001
//
//f(0b10000001) == 2
//
//f(0b11000001) == 3
//
//解法1，与1做与操作；然后进行位移操作，再与1做与操作，直到位移为0；时间复杂度O(log2(n))。

package main

import (
	"fmt"
	"strings"
)

func main() {

}

func handler1(n uint8) int {
	num := fmt.Sprintf("%b", n)
	return strings.Count(num, "1")
}

// O(n)
func handler2(n uint8) int {
	num := fmt.Sprintf("%b", n)
	var sum int
	for _, v := range num {
		if v == '1' {
			sum++
		}
	}
	return sum
}
