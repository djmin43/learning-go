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
	runesAndBytes()
}

func runesAndBytes() {
	var a rune = '0'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(a, s, b, s2)
}
