package main

import (
	"fmt"

	"github.com/djmin43/learning-go/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	slicingSlices()
}

func slicingSlices() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
}
