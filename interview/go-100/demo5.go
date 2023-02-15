package main

import "fmt"

func main() {
	fmt.Println(deferCall())
}

func deferCall() int {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	panic("4")
	//return 5
}
