package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []Users `json:"users"`
}

type User struct {
	Name 	string 		`json:"name"`
	Type 	string 		`json:"type"`
	Age 	int 			`json:"age"`
	Social Social		`json:"social"`
}

type Social struct {
	Facebook	string `json:"facebook"`
	Twitter	 	string `json:"twitter"`
}

func main() {
	fmt.Println("hello world!")

	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])
}