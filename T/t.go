package main

import (
	"fmt"
	"strconv"
)

type vestor[T any] []T
type M[K string, V any] map[K]V

type NumStr interface {
	Num | Str
}
type Num interface {
	~int | ~int32 | ~uint64
}
type Str interface {
	string
}

func add[T NumStr](a, b T) T {
	return a + b
}

func printslice[T any](s []T) { //[T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

type Price int

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

type Price2 string

func (i Price2) String() string {
	return string(i)
}

type ShowPrice interface {
	String() string
	~int | ~string
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

func main() {
	v1 := vestor[int]{58, 1881}
	printslice(v1)
	v2 := vestor[string]{"dudu", "yiyi", "8号"}
	printslice(v2)

	m1 := M[string, int]{"key": 1}
	m1["key2"] = 2

	m2 := M[string, string]{"key": "value"}
	m2["key2"] = "dudu"
	fmt.Println(m1, m2)

	fmt.Println(add(uint64(2), 4))
	fmt.Println(add("dudu", "yiyi"))

	fmt.Println(ShowPriceList([]Price{1, 2}))
	fmt.Println(ShowPriceList([]Price2{"a", "b"}))

	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(findFunc([]string{"dudu", "yiyi", "8号"}, "dudu"))
}
