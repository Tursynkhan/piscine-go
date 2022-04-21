package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for z := 0; z <= 9; z++ {
				for d := 0; d <= 9; d++ {
					if (j <= d && i < z) && (i < j) {

						z01.PrintRune('0' + rune(i))
						z01.PrintRune('0' + rune(j))
						z01.PrintRune(' ')
						z01.PrintRune('0' + rune(z))
						z01.PrintRune('0' + rune(d))
						if i == 9 && j == 8 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						continue
					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}
