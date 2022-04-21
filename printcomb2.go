package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for z := '0'; z <= '9'; z++ {
				for d := '0'; d <= '9'; d++ {
					if j < d {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(d)
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}
