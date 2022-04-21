package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					if (i <= k && j < l) || (i < k) {

						z01.PrintRune('0' + rune(i))
						z01.PrintRune('0' + rune(j))
						z01.PrintRune(' ')
						z01.PrintRune('0' + rune(k))
						z01.PrintRune('0' + rune(l))
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
	z01.PrintRune('\n')
}
