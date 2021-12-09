package main

import "fmt"

// func New() (*Database, error) {
// 	conn, err := sql.Open("postgres")
// 	if err != nil {
// 		return &Database{}, err
// 	}
// 	return &Database {
// 		Connection: conn,
// 	}, nil
// }

var name string

func init() {
	fmt.Println("This will get called on main initialization")
	name = "Jay"
}

func main() {
	fmt.Println("hello world")
	fmt.Printf("Name: %s \n", name)
}