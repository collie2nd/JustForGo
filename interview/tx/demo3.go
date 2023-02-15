package main

import "fmt"

func main() {

	a := []int{1, -2, 4, 5, -1, 1}
	fmt.Println(handler(a))
}

func handler(nums []int) int {
	if isAllMinus(nums) {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var h, r, max int
	h = 0
	r = 1
	max = maxOfInts(nums)
	for h < r {
		s := sumOfInts(nums[h:r])
		if s > max {
			max = s
		}
		if r == len(nums)-1 {
			h++
			r = h + 1
		} else {
			r++
		}
		if h == len(nums)-1 {
			break
		}
	}

	return max
}

//
//func handler(nums []int) int {
//	if isAllMinus(nums) {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	var h, r, max int
//	h = 0
//	r = 1
//	max = maxOfInts(nums)
//	for h < r {
//		s := sumOfInts(nums[h:r])
//		if s < 0 {
//			h = r + 1
//			r = h + 1
//		}
//	}
//
//	return max
//}

func isAllMinus(nums []int) bool {
	for k := range nums {
		if nums[k] > 0 {
			return false
		}
	}
	return true
}

func sumOfInts(numsTmp []int) (sum int) {
	for k := range numsTmp {
		sum += numsTmp[k]
	}
	return
}

func maxOfInts(numsTmp []int) (max int) {
	max = 0
	for k := range numsTmp {
		if numsTmp[k] > max {
			max = numsTmp[k]
		}
	}
	return
}
