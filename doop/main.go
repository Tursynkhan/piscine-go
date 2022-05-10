package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		x := 9223372036854775807
		y := -9223372036854775808
		value1 := BasicAtoi(args[0])
		value2 := BasicAtoi(args[2])
		operator := args[1]

		if value1 >= x && value2 > y || value1-x == 2 {
			return
		}
		if IsNumeric(args[0]) && IsNumeric(args[2]) {
			Calculate(value1, value2, operator)
		}

	}
}

func IsNumeric(s string) bool {
	if s[0] == '-' {
		s = s[1:]
	}
	for i := range s {
		if s[i] < 48 || s[i] > 57 {
			return false
		}
	}
	return true
}

func BasicAtoi(s string) int {
	number := 0
	negative := false

	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	if len(s) > 0 {
		for i := range s {
			number = number*10 + int(s[i]-'0')
		}
		if negative {
			number *= -1
		}
	}

	return number
}

func Calculate(a int, b int, operator string) {
	switch operator {
	case "+":
		if a == 9223372036854775807 && b > 0 || b == 9223372036854775807 && a > 0 {
			return
		}
		fmt.Println(a + b)
	case "-":
		if a == -9223372036854775808 && b > 0 || b == -9223372036854775808 && a > 0 {
			return
		}
		fmt.Println(a - b)
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(a / b)
		}
	case "*":
		if a >= 9223372036854775807/2 && b > 2 || b >= 9223372036854775807/2 && a > 2 {
			return
		} else if a <= -9223372036854775808/2 && b > 2 || b >= -9223372036854775808/2 && a > 2 {
			return
		}
		fmt.Println(a * b)
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(a % b)
		}
	default:
		return
	}
}
