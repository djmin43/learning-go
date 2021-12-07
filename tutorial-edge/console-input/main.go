package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------")

	scanner()

	// for {
	// 	fmt.Print("->")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Println(text)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	}
	// }
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
