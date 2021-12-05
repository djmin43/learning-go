package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Employee Name:" + e.Name
}

func PrintDetails (e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Jay"}
	PrintDetails(engineer)
}