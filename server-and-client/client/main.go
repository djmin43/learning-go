package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("text!!")
	resp, err := http.Post("http://localhost:18888", "text/plan", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}