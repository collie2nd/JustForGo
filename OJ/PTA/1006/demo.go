package pta

import "fmt"

func Pta1006() {
	var n int
	fmt.Scan(&n)
	//n = 12
	if n > 0 && n < 1000 {
		numberList := Digits(n)
		r := PrintInit(numberList)
		fmt.Println(string(r))
	}
}

func PrintInit(l []int) (result []rune) {
	if len(l) == 3 {
		for i := 0; i < l[0]; i++ {
			result = append(result, 'B')
		}
		for i := 0; i < l[1]; i++ {
			result = append(result, 'S')
		}
		for i := 0; i < l[2]; i++ {
			result = append(result, rune(i+49))
		}
	}
	if len(l) == 2 {
		for i := 0; i < l[0]; i++ {
			result = append(result, 'S')
		}
		for i := 0; i < l[1]; i++ {
			result = append(result, rune(i+49))
		}
	}
	if len(l) == 1 {
		for i := 0; i < l[0]; i++ {
			result = append(result, rune(i+49))
		}
	}
	return
}

func Digits(n int) (l []int) {
	var k, i, a int
	for i = 0; n >= Pow(10, i); i++ {
		k++
	}
	for i = 0; i < k; i++ {
		a = n / Pow(10, k-1-i)
		n = n - a*Pow(10, k-1-i)
		l = append(l, a)
	}
	return
}

func Pow(a, b int) (n int) {
	n = 1
	var i int
	for i = 0; i < b; i++ {
		n = n * a
	}
	return
}

