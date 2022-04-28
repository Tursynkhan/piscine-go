package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for index, value := range arguments {
		if index != 0 {
			for _, letter := range value {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
