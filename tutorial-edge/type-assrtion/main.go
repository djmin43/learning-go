package main

import (
	"fmt"
)

type Developer struct {
	Name	string
	Age		int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	return Developer{
		Name: name.(string),
		Age: age.(int),
	}
}

func main() {
	var name interface{} = "Jay"

	value, ok := name.(string)
	fmt.Println(value)
	fmt.Println(ok)

	intValue, ok := name.(int)
	fmt.Println(intValue)
	fmt.Println(ok)
}
