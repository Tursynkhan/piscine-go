package piscine

import (
	"github.com/01-edu/z01"
)

var arr = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func PrintCombN(n int) {
	recursion(0, "", n)
}

func recursion(start int, out string, n int) {
	if n == 0 {
		if out == "9" || out == "89" || out == "789" || out == "6789" || out == "56789" || out == "456789" || out == "3456789" || out == "23456789" || out == "123456789" {
			for _, v := range out {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else {
			for _, v := range out {
				z01.PrintRune(v)
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for i := start; i < 10; i++ {
		str := out + arr[i]

		recursion(i+1,
			str, n-1)

	}
}
