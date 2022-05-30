// https://leetcode.cn/problems/unique-substrings-in-wraparound-string/

package main

import "fmt"

func main() {
	fmt.Println(findSubstringInWraproundString("a"))
}

// My Code
//func findSubstringInWraproundString(p string) int {
//	dp := map[uint8][]int{}
//	dp[p[0]] = []int{1}
//	i := 1
//	for {
//		if i == len(p) {
//			break
//		}
//		if _, ok := dp[p[i]]; !ok {
//			if p[i] == p[i-1]+1 || (p[i] == 'a' && p[i-1] == 'z') {
//				dp[p[i]] = []int{dp[p[i-1]][len(dp[p[i-1]])-1] + 1}
//			} else {
//				dp[p[i]] = []int{1}
//			}
//		} else {
//			if p[i] == p[i-1]+1 || (p[i] == 'a' && p[i-1] == 'z') {
//				dp[p[i]] = append(dp[p[i]], dp[p[i-1]][len(dp[p[i-1]])-1]+1)
//			} else {
//				dp[p[i]] = append(dp[p[i]], 1)
//			}
//		}
//		i++
//	}
//	var sum int
//	for c, nums := range dp {
//		sum += getMax(nums)
//		fmt.Println(string(c), ":", nums)
//	}
//	return sum
//}
//
//func getMax(nums []int) int {
//	max := nums[0]
//	for _, n := range nums {
//		if max < n {
//			max = n
//		}
//	}
//	return max
//}

// 空间复杂度最低算法
func findSubstringInWraproundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
			k++
		} else {
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 时间复杂度最低算法
//func findSubstringInWraproundString(p string) int {
//	return method_dp(p)
//}
//
////map_abc  key = 字母x, value = 以字母x结束的个数。 如map_abc[c]=4表示 c, bc, abc, zabc
////dp[0] = 1
////dp[i] = 是相连字母 ? dp[i-1] + 1 : 1
////map_abc[p[i]] = max(dp[i], max_abc[p[i]])
////return sum(map_abc)
//func method_dp(p string) int {
//	if len(p) < 2 {
//		return len(p)
//	}
//
//	map_abc := make([]int, 123) // 用来去重
//	dp := make([]int, len(p))
//	dp[0] = 1
//	map_abc[p[0]] = 1
//
//	for i := 1; i < len(p); i++ {
//		dp[i] = 1
//		if p[i]-p[i-1] == 1 || int(p[i])-int(p[i-1]) == -25 {
//			dp[i] += dp[i-1]
//		}
//		map_abc[p[i]] = getMax(dp[i], map_abc[p[i]])
//	}
//
//	cnt := 0
//	for _, n := range map_abc {
//		cnt += n
//	}
//	return cnt
//}
//
//func getMax(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
