// https://leetcode-cn.com/problems/count-of-matches-in-tournament/

package main

import "fmt"

func main() {
	n := 7
	fmt.Println(numberOfMatches(n))
}

func numberOfMatches(n int) int {
	if n <= 1 {
		return 0
	}
	if n%2 == 1 {
		return numberOfMatches(n/2+1) + n/2
	} else {
		return numberOfMatches(n/2) + n/2
	}
}
