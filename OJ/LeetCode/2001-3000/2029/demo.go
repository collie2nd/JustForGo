// https://leetcode-cn.com/problems/stone-game-ix/

package main

import (
	"fmt"
)

func main() {
	stones := []int{5, 11, 12, 1, 4, 5, 3}
	fmt.Println(stoneGameIX(stones))
}

// 2->2/1/2/1/2/1...
// 1->1/2/1/2/1/2...

//func stoneGameIX(stones []int) bool {
//	if len(stones) == 1 {
//		return false
//	}
//
//	stoneInfos := [3]int{}
//	for _, v := range stones {
//		v = v % 3
//		if v == 0 {
//			stoneInfos[0]++
//		} else if v == 1 {
//			stoneInfos[1]++
//		} else if v == 2 {
//			stoneInfos[2]++
//		}
//	}
//
//	if (stoneInfos[1] == 0 && stoneInfos[2] <= 2) || (stoneInfos[2] == 0 && stoneInfos[1] <= 2) {
//		return false
//	}
//
//	if stoneInfos[1] == 0 || stoneInfos[2] == 0 {
//		return !(stoneInfos[0]%2 == 0)
//	}
//
//	if math.Abs(float64(stoneInfos[1]-stoneInfos[2])) <= 2 {
//		return stoneInfos[0]%2 == 0
//	}
//
//	return true
//}

// best:

func stoneGameIX(stones []int) bool {
	cnt0, cnt1, cnt2 := 0, 0, 0
	for _, val := range stones {
		val %= 3
		if val == 0 {
			cnt0++
		} else if val == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}
	if cnt0%2 == 0 {
		return cnt1 >= 1 && cnt2 >= 1
	}
	return cnt1-cnt2 > 2 || cnt2-cnt1 > 2
}
