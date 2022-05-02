package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	var runesx []rune
	var runesy []rune

	for points.x > 0 {
		runesx = append(runesx, rune((points.x%10)+48))
		points.x /= 10
	}
	for points.y > 0 {
		runesy = append(runesy, rune((points.y%10)+48))
		points.y /= 10
	}
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := len(runesx) - 1; i >= 0; i-- {
		z01.PrintRune(runesx[i])
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for j := len(runesy) - 1; j >= 0; j-- {
		z01.PrintRune(runesy[j])
	}
	z01.PrintRune('\n')
}
