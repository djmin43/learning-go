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
	maps()
}


func maps() {
	fmt.Println("maps")
}
