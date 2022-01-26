package pta

import "fmt"

var Satisfied int

func Pta1007() {
	var N int
	fmt.Scan(&N)
	//N := 4
	//fmt.Println(GetPrimeList(N))
	if N > 0 && N < 100000 {
		for i := 1; i <= N; i++ {
			if IsPrime(i) && i+2 <= N {
				if !IsPrime(i + 1) && IsPrime(i+2) {
					i ++
					Satisfied++
				}
			}
		}
		fmt.Println(Satisfied)
	}
}

func IsPrime(n int) bool {
	k := 0
	for i := 1; i <= GetRootNumber(n); i++ {
		if n%i == 0 {
			k++
		}
	}
	if k == 1 {
		return true
	}
	return false
}

func GetRootNumber(n int) int {
	i := 0
	for ; i*i <= n; i++ {
	}
	return i - 1
}

//func GetPrimeList(n int) []int {
//	var (
//		result []int
//		k      int
//	)
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= GetRootNumber(i); j++ {
//			if i%j == 0 {
//				k++
//			}
//		}
//		if k == 1 {
//			result = append(result, i)
//		}
//		k = 0
//	}
//	return result
//}
