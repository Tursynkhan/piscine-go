package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, value := range a {
		if value == " " {
			z01.PrintRune('\n')
		}
	}
}
