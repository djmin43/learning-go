package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello world")
	// fmt.Println(mascot.BestMascot())
	// fmt.Println(quote.Go())
	// fmt.Println("Greetings and \n\"Salutations\"")
	isStringLarger := "a" > "b"
	fmt.Println(isStringLarger)
	fmt.Println(typeConversion())
}

func typeConversion () float64{
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	return z
}
