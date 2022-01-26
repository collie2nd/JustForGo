package pta

import (
	"fmt"
	"sort"
)

func Pta1001Plus() {
	var K int
	fmt.Scanln(&K)
	numberList := make([]int, K)
	pathList := make([][]int, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&numberList[i])
	}

	//numberList := []int{2, 4, 5}
	//pathList := make([][]int, 3)
	for k, v := range numberList {
		if v > 1 && v <= 100 {
			for {
				if v%2 == 1 {
					v = (3*v + 1) / 2
				} else {
					v = v / 2
				}
				if v == 1 {
					break
				}
				pathList[k] = append(pathList[k], v)
			}
		}
	}
	fmt.Println(pathList)
	var (
		i                       int
		keyIndexList, indexList []int
		keyList                 [][]int
	)
	keyList = append(keyList, pathList[0])
	keyIndexList = []int{0}
	for k, v := range pathList {
		//flag1:是否包含该路径，flag2：是否路径相同
		flag1, flag2 := IsAppend(&v, &keyList[i])
		if !flag1 {
			if flag2 {
				keyIndexList[i] = k
			} else {
				keyList = append(keyList, v)
				keyIndexList = append(keyIndexList, k)
				i++
			}
		}
	}
	for _, v := range keyIndexList {
		indexList = append(indexList, numberList[v])
	}
	sort.Ints(indexList)
	result := []int{indexList[0]}
	for i := 0; i+1 < len(indexList); i++ {
		if indexList[i] == indexList[i+1] {
			continue
		}
		result = append(result, indexList[i+1])
	}
	for i := len(result) - 1; i >= 0; i-- {
		if i != 0 {
			fmt.Printf("%v ", result[i])
		} else {
			fmt.Printf("%v", result[i])
		}
	}
}

//bool1:是否包含该路径，bool2：是否路径相同
func IsAppend(a, path *[]int) (bool, bool) {
	if len(*a) <= len(*path) {
		for i := 1; i < len(*a); i++ {
			if (*a)[len(*a)-i] != (*path)[len(*path)-i] {
				return false, false
			}
		}
		return true, true
	} else {
		for i := 1; i < len(*a); i++ {
			if (*a)[len(*a)-i] != (*path)[len(*path)-i] {
				return false, false
			}
			if len(*path) == i {
				break
			}
		}
		*path = *a
		return false, true
	}
}
