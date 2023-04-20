package main

import "fmt"

func main() {
	a := []int{1, -2, 3, -1, 4, 2, -3, 4, 7}
	fmt.Println(getMaxSubSlice(a))
}
func getMaxSubSlice(src []int) (rst []int) {
	var i, j int
	for i < j {
		if j == len(src) {
			break
		}

	}
	return
}

func getSum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
