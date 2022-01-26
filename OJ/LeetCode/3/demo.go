// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/submissions/

package main

import (
	"fmt"
)

func main() {
	//var s string
	//fmt.Scan(&s)
	s := "abcdcda"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s) - 0
	}

	isExist := make(map[byte]int)
	maxLen := 0
	for k, v := range s {
		if index, exist := isExist[byte(v)]; exist {
			isExist = map[byte]int{}
			for i := index + 1; i < k; i++ {
				isExist[s[i]] = i
			}
		}
		isExist[byte(v)] = k
		if maxLen < len(isExist) {
			maxLen = len(isExist)
		}
	}
	return maxLen
}

//空间复杂度最低算法：
//func lengthOfLongestSubstring(s string) (ret int) {
//	l, cache := 0, make([]int, 128)
//	for i := 0; i < len(s); i++ {
//		idx := s[i]
//		l = int(math.Max(float64(l), float64(cache[idx])))
//		ret = int(math.Max(float64(ret), float64(i-l+1)))
//		cache[idx] = i + 1
//	}
//	return
//}

//时间复杂度最低算法：
//func lengthOfLongestSubstring(s string) int {
//	if len(s) <= 1 {
//		return len(s) - 0
//	}
//	res := 0
//	//字母表
//	letterMap := make(map[string]int, len(s))
//	left, right := 0, 0 //左右指针
//	for ; right < len(s); right++ {
//		var letter string
//		//当前位置的字母
//		if right == len(s)-1 { //最后一个字符
//			letter = s[right:]
//		} else {
//			letter = s[right : right+1]
//		}
//
//		_, ok := letterMap[letter]
//		letterMap[letter] += 1
//		if !ok {
//			continue
//		}
//		//出现重复，将当前最长串记录
//		size := len(letterMap)
//		if size > res {
//			res = size
//		}
//		//缩小左边边界，直到字母不重复
//		for {
//			if letterMap[letter] <= 1 && left < len(s) {
//				break
//			}
//			var str string
//			if left == len(s)-1 {
//				str = s[left:]
//			} else {
//				str = s[left : left+1]
//			}
//			if str == letter {
//				letterMap[letter] -= 1
//			} else {
//				delete(letterMap, str)
//			}
//			left++
//		}
//	}
//	//最后记得和map大小进行对比,因为可能存在整个字符串都是不相同字符的
//	if len(letterMap) > res {
//		res = len(letterMap)
//	}
//	return res
//}
