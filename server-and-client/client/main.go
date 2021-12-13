package main

import "mime/multipart"

func main() {
	var buffer bytes.buffer
	writer := multipart.NewWriter(&buffer)
	
}