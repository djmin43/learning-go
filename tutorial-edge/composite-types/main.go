package main

import "fmt"

func array () {
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days)
}

func slices () {
	// Slices are made up of three things, a pointer, a length, and a capacity. 
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays := days[0:5]
	fmt.Println(weekdays)
}

func maps () {
	youtubeSubscribers := map[string]string{
		"TutorialEdge":     "hello",
		"MKBHD":            "world",
		"Fun Fun Function": "my name is",
	}
	fmt.Println(youtubeSubscribers["MKBHD"])	
}

func main () {
	array()
	slices()
	maps()
}