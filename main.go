package main

import "fmt"

func main() {
	twoBase := makeMult(2)
	fmt.Println(twoBase(2))
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * base
	}
}
