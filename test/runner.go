package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func Atoi(s string)int{

// }

// func Expandstr(s string) string {
// 	var result []string
// 	word := ""
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != ' ' {
// 			word += string(s[i])
// 			if i+1 == len(s) {
// 				result = append(result, word)
// 			}
// 		} else {
// 			if word != "" {
// 				result = append(result, word)
// 			}
// 			word = ""
// 		}
// 	}
// 	answer := ""
// 	for i := 0; i < len(result); i++ {
// 		if i+1 == len(result) {
// 			answer += result[i]
// 			break
// 		}
// 		answer += result[i] + "   "
// 	}
// 	return answer
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	fmt.Println(os.Args[1])
// 	fmt.Println(Expandstr(os.Args[1]))
// }

// func ReverseBits(oct byte) byte {
// 	res := 0
// 	num := int(oct)
// 	for i := 0; i < 8; i++ {
// 		res = res*2 + num%2
// 		num /= 2
// 	}
// 	return byte(res)
// }

// func main() {
// 	fmt.Println(ReverseBits(16))
// }
func RepeatAlpha(s string) {
out:
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			for j := 0; j < CountIt(rune(s[i]), true); j++ {
				z01.PrintRune(rune(s[i]))
			}
			continue out
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			for j := 0; j < CountIt(rune(s[i]), false); j++ {
				z01.PrintRune(rune(s[i]))
			}
			continue out
		}
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}

func CountIt(s rune, signal bool) int {
	n := 0
	if signal {
		n = int(s - 96)
	} else {
		n = int(s - 64)
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	RepeatAlpha(os.Args[1])
}
