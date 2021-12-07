package pta

import (
"fmt"
)

func PATPASS() {
	var n int
	fmt.Scanf("%v", &n)
	test := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &test[i])
	}
	//test := []string{
	//	"APAAATAA",
	//}

	for _, v := range test {
		if TestBase(v) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

//index标注着PAT的首位元素下标
func TestBase(s string) bool {
	arr := []rune(s)
	if len(arr) > 3 {
		for _, v := range arr {
			if v != 'P' && v != 'A' && v != 'T' {
				return false
			}
		}
		n := len(arr)
		var indexForP, indexForAT int
		for i := 0; i <= n; i++ {
			if arr[i] == 'P' {
				indexForP = i
				break
			}
		}
		for i := 0; i+2 <= n; i++ {
			if s[i:i+2] == "AT" {
				indexForAT = i
				break
			}
		}
		if indexForAT != 0 {
			if indexForP+1 != indexForAT {
				a := arr[:indexForP]
				b := arr[indexForP+1 : indexForAT]
				ca := arr[indexForAT+2:]
				c := arr[indexForAT+2 : indexForAT+2+len(ca)-len(a)]
				if TestString(a) && TestString(b) && TestString(c) &&len(c) == len(a)*len(b){
					//if len(c) > len(a) {
					//	if len(c) == len(a)*len(b) {
					//		return true
					//	}
					//} else if len(c) == len(a) {
					//	if len(b) == 1 {
					return true
					//	}
					//}
				}
			} else {
				if indexForP == len(s)-indexForAT-2 {
					return true
				}
			}
		}
	} else if s == "PAT" {
		return true
	}
	return false
}

func TestString(s []rune) bool {
	if s != nil {
		for _, v := range s {
			if v != 'A' {
				return false
			}
		}
	}
	return true
}

//func Pow(a, b int) (n int) {
//	n = 1
//	var i int
//	for i = 0; i < b; i++ {
//		n = n * a
//	}
//	return
//}

