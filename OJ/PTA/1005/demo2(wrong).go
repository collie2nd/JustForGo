package pta

import (
	"fmt"
	"sort"
)

type Path struct {
	PathSource int
	PathInfo   *[]int
}

type NumberInfo struct {
	Number  int
	PathTo1 []int
}

func Pta1005() {
	var K int
	fmt.Scan(&K)
	NumList := make([]NumberInfo, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&NumList[i].Number)
	}

	//NumList := []NumberInfo{
	//	{
	//		3,
	//		[]int{},
	//	},
	//	{
	//		5,
	//		[]int{},
	//	},
	//	{
	//		6,
	//		[]int{},
	//	},
	//	{
	//		7,
	//		[]int{},
	//	},
	//	{
	//		8,
	//		[]int{},
	//	},
	//	{
	//		11,
	//		[]int{},
	//	},
	//}

	for i := 0; i < len(NumList); i++ {
		num := NumList[i].Number
		if num > 1 && num <= 100 {
			for {
				if num%2 == 1 {
					num = (3*num + 1) / 2
				} else {
					num = num / 2
				}
				if num == 1 {
					break
				}
				NumList[i].PathTo1 = append(NumList[i].PathTo1, num)
			}
		}
	}
	pathInit := &NumList[0].PathTo1
	PathList := []Path{{0, pathInit,}}
	//PathList := []Path{}
	//找到参照的Path
	//比对每个数字对应的PATH
	for i := 0; i < len(NumList); i++ {
		for j := 0; j < len(PathList); j++ {
			b := *PathList[j].PathInfo
			basePathLength := len(b)
			t := NumList[i].PathTo1
			thisNumPathLength := len(t)
			//如果参照路径比当前路径长，则看是否当前路径完全从属于参照路径，如果属于则非独有路径，如果不属于则属于独有路径
			if basePathLength >= thisNumPathLength {
				if !SliceEqual(b[basePathLength-thisNumPathLength:], t[:]) {
					PathList = append(PathList, Path{})
					PathList[j+1].PathInfo = &NumList[i].PathTo1
					PathList[j+1].PathSource = i
					break
				}
			} else {
				//如果参照路径比当前路径短，则看是否参照路径属于当前路径，如果属于可以更新当前参照路径为当前路径，如果不属于，则当前路径为独有路径
				if SliceEqual(b[:], t[thisNumPathLength-basePathLength:]) {
					PathList[j].PathInfo = &NumList[i].PathTo1
					PathList[j].PathSource = i
				} else {
					if j == len(PathList)-1 {
						PathList = append(PathList, Path{})
						PathList[j+1].PathInfo = &NumList[i].PathTo1
						PathList[j+1].PathSource = i
						break
					}
				}
			}
		}
	}
	//fmt.Println(NumList)
	result := []int{}
	for i := 0; i < len(PathList); i++ {
		if PathList[i].PathInfo != nil {
			result = append(result, NumList[PathList[i].PathSource].Number)
		}
	}
	sort.Ints(result)
	for i := len(result) - 1; i >= 0; i-- {
		if i != len(result)-1 {
			fmt.Printf(" %v", result[i])
		} else {
			fmt.Print(result[i])
		}
	}
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
}
