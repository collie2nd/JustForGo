package main

import (
	"fmt"
	"strconv"
	"strings"
)

func transfer(num int) string {
	chineseMap := []string{"圆整", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	var listNum []int
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}
	n := len(listNum)
	chinese := ""
	// 注意这里是倒序的，注意替换顺序
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	for {
		copyChinese := chinese
		copyChinese = strings.Replace(copyChinese, "零万", "万", 1)
		copyChinese = strings.Replace(copyChinese, "零亿", "亿", 1)
		copyChinese = strings.Replace(copyChinese, "零十", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零百", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零千", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零零", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零圆", "圆", 1)

		if copyChinese == chinese {
			break
		} else {
			chinese = copyChinese
		}
	}

	return chinese
}

func main() {
	fmt.Println(1024, transfer(1024))
	fmt.Println(1004, transfer(1004))
	fmt.Println(101004, transfer(101024))
	fmt.Println(1010004, transfer(1010004))
	fmt.Println(3000010004, transfer(3000010004))
	fmt.Println(3000010004.1, ConvertNumToCny(3000010004.1))
}

func ConvertNumToCny(num float64) string {
	strNum := strconv.FormatFloat(num*1000, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分", "离"}
	// log.Println(sliceUnit[:len(sliceUnit)-2])
	s := sliceUnit[len(sliceUnit)-len(strNum):]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	chinese := ""
	for k, v := range strNum[:] {
		chinese = chinese + upperDigitUnit[string(v)] + s[k]
	}
	for {
		copyChinese := chinese
		copyChinese = strings.Replace(copyChinese, "零万", "万", 1)
		copyChinese = strings.Replace(copyChinese, "零亿", "亿", 1)
		copyChinese = strings.Replace(copyChinese, "零拾", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零佰", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零仟", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零零", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零圆", "圆", 1)

		if copyChinese == chinese {
			break
		} else {
			chinese = copyChinese
		}
	}
	for {
		copyChinese := chinese
		copyChinese = strings.Replace(copyChinese, "零零", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零角", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零分", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零离", "", 1)

		if copyChinese == chinese {
			break
		} else {
			chinese = copyChinese
		}
	}
	// 中文乱码问题
	if chinese[len(chinese)-3:] == "圆零" {
		chinese = strings.Replace(chinese, "圆零", "圆整", 1)
	}
	if chinese[len(chinese)-2:] == "零" {
		chinese = strings.Replace(chinese, "零", "", 1)
	}

	return chinese
}
