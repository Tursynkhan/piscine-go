package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) int {
	if x < 1 || y < 1 {
		return 0
	} else {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('/')
					} else if j == x {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('\\')
					} else if j == x {
						z01.PrintRune('/')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			} else if i > 1 && i < y {
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}
