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
	julia := person{
		"julia",
		40,
		"cat",
	}
	sang := person{
		name: "sang",
		age:  100,
		pet:  "vayne",
	}
	fmt.Println(sang)
	fmt.Println(julia)
}
