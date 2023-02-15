package main

import (
	"fmt"
)

func main() {
	var n, T int
	var jobs [][2]int
	fmt.Scan(&T, &n)
	for i := 0; i < n; i++ {
		var t, w int
		fmt.Scan(&t, &w)
		jobs = append(jobs, [2]int{t, w})
	}
	fmt.Println(findBestChoice(jobs, T))
}

func findBestChoice(jobsTmp [][2]int, t int) int {
	if len(jobsTmp) == 0 {
		return 0
	}
	if t <= 0 {
		return 0
	}
	if len(jobsTmp) == 1 {
		if t >= jobsTmp[0][0] {
			return jobsTmp[0][1]
		} else {
			return 0
		}
	}

	var max int
	for k, v := range jobsTmp {
		jobsChildTmp := append(jobsTmp[:k], jobsTmp[k+1:]...)
		tmp := v[1] + findBestChoice(jobsChildTmp, t-v[0])
		if tmp > max {
			max = tmp
		}
	}
	return max
}
