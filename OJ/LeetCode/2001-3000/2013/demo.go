// https://leetcode-cn.com/problems/...

package main

import (
	"fmt"
)

func main() {
	operations := []string{"DetectSquares", "add", "add", "add", "count", "count", "add", "count"}
	points := [][]int{{}, {3, 10}, {11, 2}, {3, 2}, {11, 10}, {14, 8}, {11, 2}, {11, 10}}

	obj := Constructor()

	for i := 1; i < len(operations); i++ {
		point := points[i]
		operation := operations[i]

		if operation == "add" {
			obj.Add(point)
		}

		if operation == "count" {
			param2 := obj.Count(point)
			fmt.Println(param2)
		}
	}

}

// My Code
type DetectSquares map[int]map[int]int

func Constructor() DetectSquares {
	return DetectSquares{}
}

func (s DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if s[y] == nil {
		s[y] = map[int]int{}
	}
	s[y][x]++
}

func (s DetectSquares) Count(point []int) (ans int) {
	x, y := point[0], point[1]
	if s[y] == nil {
		return
	}
	yCnt := s[y]
	for col, colCnt := range s {
		if col != y {
			// 根据对称性，这里可以不用取绝对值
			d := col - y
			ans += colCnt[x] * yCnt[x+d] * colCnt[x+d]
			ans += colCnt[x] * yCnt[x-d] * colCnt[x-d]
		}
	}
	return
}

// 空间复杂度最低算法

// 时间复杂度最低算法
//const (
//	N = 1001
//)
//
//type DetectSquares struct {
//	m map[int]int
//	l [1001][]int
//}
//
//func ref(x, y int) int {
//	return x*N+y
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func Constructor() DetectSquares {
//	return DetectSquares{
//		m: make(map[int]int, N),
//		l: [1001][]int{},
//	}
//}
//
//
//func (t *DetectSquares) Add(p []int)  {
//	f := ref(p[0], p[1])
//	t.m[f] ++
//	if t.m[f] == 1 {
//		t.l[p[0]] = append(t.l[p[0]], p[1])
//	}
//}
//
//
//func (t *DetectSquares) Count(p []int) int {
//	ret := 0
//	for _, y := range t.l[p[0]] {
//		if y == p[1] {
//			continue
//		}
//		ll := abs(y-p[1])
//		bas := t.m[ref(p[0], y)]
//		if p[0]+ll >= 0 && p[0]+ll < N {
//			ret += bas * t.m[ref(p[0]+ll, p[1])] * t.m[ref(p[0]+ll, y)]
//		}
//		if p[0]-ll >= 0 && p[0]-ll < N {
//			ret += bas * t.m[ref(p[0]-ll, p[1])] * t.m[ref(p[0]-ll, y)]
//		}
//	}
//	return ret
//}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
