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
	declareSlices()
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

func declareSlices() {
	var a = []int{10, 20, 30, 40}
	fmt.Println(a)
	// var b [][]int
	var c []int
	fmt.Println(c)
	// nil is an identifier that represents the lack of a value for some types
	// nil has no type
	// slices cannot be compared
	fmt.Println(c == nil)
	var d []int
	d = append(d, 10)
	fmt.Println(d)

	e := make([]int, 5)
	fmt.Println(e)
	f := make([]int, 5, 10)
	fmt.Println(f)

	// when we create a slice, the primary goal is to minimize the number of times the slice needs to grow.
	// slice literal is a good choice if values are not going to change
	// if not sure, use make.
}
