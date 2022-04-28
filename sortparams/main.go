package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for j := 1; j < len(arg); j++ {
		for _, element := range arg[j] {
			z01.PrintRune(element)
		}
		z01.PrintRune('\n')
	}
}
