package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	name := os.Args[1:]
	if len(name) == 0 {
		fmt.Println("File name missing")
	} else if len(name) > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, err := ioutil.ReadFile(string(name[0]))
		if err != nil {
			fmt.Println(err.Error)
		}
		fmt.Println(string(file[:len(file)-1]))
	}
}
