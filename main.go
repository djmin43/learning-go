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
	var s string = "hello, go"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
	// byte is 8 bit
	// rune is 32 bit (4 byte)
	// UTF-8 can be anywhere from one to four bytes long
}
