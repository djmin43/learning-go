package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	out, err := exec.Command("ls").Output()


	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output = string(out[:])
	fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("CAn't execute this on a window machine") 
	} else {
		execute()
	}
}