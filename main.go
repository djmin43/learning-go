package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if n := rand.Intn(10); n == 0 {
		fmt.Println("that is too low")
	} else if n > 5 {
		fmt.Println("that is too big", n)
	} else {
		fmt.Println("that is a good number", n)
	}
}
