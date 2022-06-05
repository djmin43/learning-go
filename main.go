package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}
	var fred person
	fmt.Println(fred)
	bob := person{}
	fmt.Println(bob)
}
