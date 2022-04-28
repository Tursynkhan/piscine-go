package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	runes := arguments[0]

	for _, value := range runes {
		if value != '.' && value != '/' {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}
