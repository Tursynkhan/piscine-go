package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aStr := []rune(s)
	for i := 0; i < len(aStr); i++ {
		z01.PrintRune(rune(aStr[i]))
	}
}
