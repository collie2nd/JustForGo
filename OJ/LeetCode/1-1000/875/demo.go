// https://leetcode.cn/problems/koko-eating-bananas/submissions/

package main

import "fmt"

func main() {

	fmt.Println()
}

// My Code

// 空间复杂度最低算法
func function(piles []int, x int) int {
	h := 0
	for i := 0; i < len(piles); i++ {
		h += piles[i] / x
		if piles[i]%x > 0 {
			h++
		}
	}
	return h
}

//func minEatingSpeed(piles []int, h int) int {
//	left := 1
//	right := 1000000000
//	for {
//		if left > right {
//			break
//		}
//		mid := (left + right) / 2
//		if function(piles, mid) == h {
//			right = mid - 1
//		} else if function(piles, mid) > h {
//			left = mid + 1
//		} else if function(piles, mid) < h {
//			right = mid - 1
//		}
//	}
//	return left
//}

// 时间复杂度最低算法
func minEatingSpeed(piles []int, h int) int {
	min, max := 1, 1<<31
	// for i := 0; i < len(piles); i++ {
	//     if piles[i] > max {
	//         max = piles[i]
	//     }
	//     if piles[i] < min {
	//         min = piles[i]
	//     }
	// }

	for min <= max {
		mid := (min + max) >> 1
		t := timeConsuming(piles, mid)
		if t == h {
			max = mid - 1
		} else if t > h {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return min
}

func timeConsuming(piles []int, speed int) int {
	ans := 0
	for i := 0; i < len(piles); i++ {
		ans += cell(float64(piles[i]) / float64(speed))
	}
	return ans
}

func cell(a float64) int {
	if a > float64(int(a)) {
		return int(a) + 1
	}
	return int(a)
}
