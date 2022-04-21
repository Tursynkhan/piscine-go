package piscine

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	a := 13
	b := 2
	div := a / b
	mod := a % b
	piscine.DivMod(a, b, &div, &mod)
	fmt.PrintLn(div)
	fmt.PrintLn(mod)
}
