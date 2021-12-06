package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}