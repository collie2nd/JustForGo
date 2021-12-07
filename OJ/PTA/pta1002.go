package pta

import (
	"fmt"
	"strings"
)

func WriteThisNumber() {

	var (
		n     []byte
		m     uint64
		count []string
		py    string
	)
	fmt.Scanf("%v", &n)
	if n[0] != '-' {
		//a := Digits(n)
		for _, v := range n {
			v -= 48
			m += uint64(v)
		}
		b := Digits(m)
		for _, v := range b {
			count = append(count, Convert(v))
		}
		py = strings.Join(count, " ")
		fmt.Println(py)
	}
}

func Digits(n uint64) (l []uint64) {
	var k, i, a uint64
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

func Convert(n uint64) (cn string) {
	if n == 1 {
		cn = "yi"
	} else if n == 2 {
		cn = "er"
	} else if n == 3 {
		cn = "san"
	} else if n == 4 {
		cn = "si"
	} else if n == 5 {
		cn = "wu"
	} else if n == 6 {
		cn = "liu"
	} else if n == 7 {
		cn = "qi"
	} else if n == 8 {
		cn = "ba"
	} else if n == 9 {
		cn = "jiu"
	} else if n == 0 {
		cn = "ling"
	}
	return
}

func Pow(a, b uint64) (n uint64) {
	n = 1
	var i uint64
	for i = 0; i < b; i++ {
		n = n * a
	}
	return
}
