// https://leetcode-cn.com/problems/two-sum/

package main

import (
	"fmt"
)

func main() {
	nums := []int{150,24,79,50,88,345,3}
	target := 200
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	history := make(map[int]int)
	for i, v := range nums {
		if res, ok := history[target - v]; ok {
			return []int{res, i}
		}
		history[v] = i
	}
	return nil
}

//func twoSum(nums []int, target int) []int {
//	numsCache := make([]int, 0)
//	rst := make([]int, 2)
//	restNums := make([]int, 0)
//	for _, v := range nums {
//		numsCache = append(numsCache, v)
//	}
//	for i := 0; i < len(numsCache); i++ {
//		restNums = append(numsCache[:i], numsCache[i+1:]...)
//		sort.Ints(restNums)
//		if findRest(restNums, target-nums[i]) {
//			rst[0] = i
//			rst[1] = target - nums[i]
//			break
//		}
//		numsCache = make([]int, 0)
//		for _, v := range nums {
//			numsCache = append(numsCache, v)
//		}
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == rst[1] && i != rst[0] {
//			rst[1] = i
//			break
//		}
//	}
//	return rst
//}
//
//func findRest(restNums []int, target int) bool {
//	if len(restNums) == 1 {
//		return restNums[0] == target
//	} else {
//		if target > restNums[len(restNums)/2] {
//			return findRest(restNums[len(restNums)/2:], target)
//		} else if target < restNums[len(restNums)/2] {
//			return findRest(restNums[0:len(restNums)/2], target)
//		} else {
//			return true
//		}
//	}
//}
