package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values {
		"test": {"hello world"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}