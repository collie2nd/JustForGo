package main

import (
	"fmt"
)

func main() {
	letterCh := make(chan struct{})
	numCh := make(chan struct{})
	endCh := make(chan struct{})
	go func() {
		for i := 'A'; i < 'Z'; i = i + 2 {
			select {
			case <-letterCh:
				fmt.Printf("%c%c", i, i+1)
				numCh <- struct{}{}
			}
		}
	}()
	go func() {
		for i := 1; ; i = i + 2 {
			select {
			case <-numCh:
				fmt.Printf("%d%d", i, i+1)
				if i == 27 {
					endCh <- struct{}{}
					return
				}
				letterCh <- struct{}{}
			}
		}
	}()
	numCh <- struct{}{}
	select {
	case <-endCh:
		return
	}
}
