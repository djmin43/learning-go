package main

import "fmt"

func main() {

	uniqueNames := map[string]bool{"fred": true, "raul": true, "wilma": true}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}
}
