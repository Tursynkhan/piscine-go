package main

import (
	"os"

	"github.com/01-edu/z01"
)

var (
	OddMsg  string = "I have an odd number of arguments"
	EvenMsg string = "I have an even number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	lengthOfArg := len(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
