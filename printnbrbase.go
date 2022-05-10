package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == -9223372036854775808 {
		str := "-9223372036854775808"
		for _, i := range str {
			z01.PrintRune(rune(i))
		}
		return
	}
	k := []rune(base)
	for i := 0; i < len(k); i++ {
		for j := 0; j < len(k); j++ {
			if k[i] == k[j] && i != j {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		if k[i] == '-' || k[i] == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	var bnm []rune
	for nbr > 0 {
		bnm = append(bnm, k[nbr%len(k)])
		nbr /= len(k)
	}
	for i := len(bnm) - 1; i >= 0; i-- {
		z01.PrintRune(bnm[i])
	}
}
