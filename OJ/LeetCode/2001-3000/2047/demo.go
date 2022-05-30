// https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//sentence := "cat and  dog"
	//sentence := "!this  1-s b8d!"
	//sentence := "alice and  bob are playing stone-game10"
	sentence := "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."
	fmt.Println(countValidWords(sentence))
}

const letters = "abcdefghijklmnopqrstuvwxyz"
const nums = "0123456789"
const punctuations = "!.,"
const hyphen = "-"

//func countValidWords(sentence string) int {
//	strs := strings.Fields(sentence)
//	sumWords := 0
//	for _, str := range strs {
//		containNum := false
//		containHyphen := []int{}
//		containPunctuation := []int{}
//		for k, v := range str {
//			if strings.Contains(nums, string(v)) {
//				containNum = true
//				break
//			}
//			if strings.Contains(hyphen, string(v)) {
//				containHyphen = append(containHyphen, k)
//			}
//			if strings.Contains(punctuations, string(v)) {
//				containPunctuation = append(containPunctuation, k)
//			}
//		}
//		if containNum {
//			continue
//		}
//		if len(containHyphen) > 1 {
//			continue
//		} else if len(containHyphen) == 1 {
//			index := containHyphen[0]
//			if index == 0 || index == len(str)-1 {
//				continue
//			}
//			if !strings.Contains(letters, str[index+1:index+2]) || !strings.Contains(letters, str[index-1:index]) {
//				continue
//			}
//		}
//		if len(containPunctuation) > 1 {
//			continue
//		} else if len(containPunctuation) == 1 {
//			if containPunctuation[0] != len(str)-1 {
//				continue
//			}
//		}
//		fmt.Println(str)
//		sumWords++
//	}
//	return sumWords
//}


func valid(s string) bool {
	hasHyphens := false
	for i, ch := range s {
		if unicode.IsDigit(ch) || strings.ContainsRune("!.,", ch) && i < len(s)-1 {
			return false
		}
		if ch == '-' {
			if hasHyphens || i == 0 || i == len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
				return false
			}
			hasHyphens = true
		}
	}
	return true
}

func countValidWords(sentence string) (ans int) {
	for _, s := range strings.Fields(sentence) { // 按照空格分割
		if valid(s) {
			ans++
		}
	}
	return
}
