package main

import (
	"fmt"
	"sort"
)

func RemoveDuplicate(v []int) []int {
	// 为了性能需要尽可能的减小拷贝，最悲观的情况每个元素只移动一次。
	toIndex := 0
	p := 0

	for i := range v {
		// 为了实际去重结构时减小内存拷贝
		c := &v[i]

		if p == *c && i != 0 {
			// 重复内容，跳过
			continue
		}

		if i != toIndex {
			// 需要移动当前元素
			v[toIndex] = *c
		}

		toIndex++
		p = *c
	}

	return v[:toIndex]
}

func main() {
	v := []int{9, 1, 1, 9, 2, 2, 3, 3, 3, 4, 5, 6, 7, 7, 7, 7, 8}

	// 升序排序
	//sort.Slice(v, func(i, j int) bool { return v[i] <= v[j] })
	sort.Ints(v)

	// 去重
	v = RemoveDuplicate(v)

	// 打印结果
	// []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v", v)
}
