// https://leetcode-cn.com/problems/single-element-in-a-sorted-array

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(a))
}

func singleNonDuplicate(a []int) int {
	i := sort.Search(len(a)-1, func(i int) bool {
		return a[i] != a[i^1]
	})
	return a[i]
}
