package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i := (len(arguments) - 1); i > 0; i-- {
		if i != 0 {
			for _, letter := range arguments[i] {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
