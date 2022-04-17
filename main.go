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
<<<<<<< HEAD
	slicingSlices()
}

func slicingSlices() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
}


=======
}

func slicingSlices() {

}
>>>>>>> 124a67cd855e212be3a14983c075fd7c84f18923
