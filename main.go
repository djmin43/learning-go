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
	declareSlices()
}

func declareSlices() {

	f := make([]int, 0, 10)
	fmt.Println(f)

	// when we create a slice, the primary goal is to minimize the number of times the slice needs to grow.
	// slice literal is a good choice if values are not going to change
	// if not sure, use make.
}
