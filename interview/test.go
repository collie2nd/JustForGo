package main

import "fmt"

//	func main() {
//		var m sync.Map
//		m.Store("address", map[string]string{"province": "江苏", "city": "南京"})
//		v, _ := m.Load("address")
//		fmt.Printf("%+v\n", v)
//	}

func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}
}
