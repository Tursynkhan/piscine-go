package main

import (
	"fmt"
)

func Concat(str1 string, str2 string) string {
	var result string
	result = str1 + str2
	return result
}

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}
