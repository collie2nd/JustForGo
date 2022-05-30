// https://leetcode.cn/problems/validate-ip-address/

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(validIPAddress("2001:0db8:85a3:00000:0:8A2E:0370:7334"))
}

// My Code
//func validIPAddress(queryIP string) string {
//	s := []string{}
//	for i := 0; i < len(queryIP); i++ {
//		switch queryIP[i] {
//		case '.':
//			s = strings.Split(queryIP, ".")
//			for _, str := range s {
//				if len(str) == 0 || len(str) > 4 {
//					return "Neither"
//				}
//			}
//		case ':':
//			s = strings.Split(queryIP, ":")
//			for _, str := range s {
//				if len(str) == 0 || len(str) > 4 {
//					return "Neither"
//				}
//			}
//		}
//	}
//	if net.ParseIP(queryIP) != nil {
//		for i := 0; i < len(queryIP); i++ {
//			switch queryIP[i] {
//			case '.':
//				return "IPv4"
//			case ':':
//				return "IPv6"
//			}
//		}
//
//	}
//	return "Neither"
//}

// 空间复杂度最低算法
func validIPAddress(queryIP string) string {

	/*
		1. 验证是否合法ipv4地址
	*/
	checkIPv4 := func(str string) bool {
		if len(str) == 0 || len(str) > 3 { //最多3位数255
			return false
		}
		num := 0
		for i, ch := range str {
			if ch < '0' || ch > '9' {
				return false
			}
			if ch == '0' && i == 0 && len(str) > 1 {
				return false
			}
			num = num*10 + int(ch-'0')
		}
		if num > 255 {
			return false
		}
		return true
	}

	checkIPv6 := func(str string) bool {
		if len(str) == 0 || len(str) > 4 { //最多4位数
			return false
		}
		for _, ch := range str {
			if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')) {
				return false
			}
		}
		return true
	}

	//先尝试判断是个ipv4
	strs := strings.Split(queryIP, ".")
	if len(strs) == 4 {
		valid := true
		for i := range strs {
			if !checkIPv4(strs[i]) {
				valid = false
				break
			}
		}
		if valid {
			return "IPv4"
		}
	}

	//再尝试判断是个ipv6
	strs = strings.Split(queryIP, ":")
	if len(strs) == 8 {
		valid := true
		for i := range strs {
			if !checkIPv6(strs[i]) {
				valid = false
				break
			}
		}
		if valid {
			return "IPv6"
		}
	}

	return "Neither"
}

// 时间复杂度最低算法
