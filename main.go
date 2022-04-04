package main

import (
	"fmt"

	"github.com/djmin43/learning-go/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	declareArray()
}

func declareArray() {
	var x [3]int
	var y = [3]int{10, 20, 30}
	var a = [...]int{1, 2, 3}
	var b [2][3]int
	fmt.Println(a, b, x, y)
	// arrays in Go has different types per sizes. Go also can't use type conversion to convert arrays of different sizes to identical types.
	// It is almost impossible to use array in Go unless you know the size of the array for sure.
}
