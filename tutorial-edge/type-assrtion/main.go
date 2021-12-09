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
	fmt.Println("hello world")

	var name interface{} = "Jay"
	var age interface{} = 33

	newDev := GetDeveloper(name, age)
	fmt.Println(newDev.Name)
	fmt.Println(newDev.Age)

}