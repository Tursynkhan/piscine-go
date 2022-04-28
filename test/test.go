package main

import (
	"fmt"
)

func BasicJoin(elems []string) string {
	result := ""
	for _, value := range elems {
		result += value
	}
	return result
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
