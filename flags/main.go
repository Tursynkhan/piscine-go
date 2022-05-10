package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1:]
	flag(name)
}

func flag(name []string) {
	if len(name) == 0 {
		help()
		return
	}
	var number string
	var str string
	var order_true bool = false
	lenstr := len(name)
	for _, w := range name {
		if len(insert(w)) > 0 {
			number += insert(w)
			lenstr--
		} else if w == "--order" || w == "-o" {
			order_true = true
			lenstr--
		} else if w == "--help" || w == "-h" || lenstr == 0 {
			help()
			return
		} else {
			str += w
		}
	}
	if len(str) >= 0 {
		str += number
		if order_true {
			fmt.Println(order(str))
		} else {
			fmt.Println(str)
		}
	} else {
		help()
	}
}

func help() {
	fmt.Println("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func order(s string) string {
	k := []rune(s)
	for i := 1; i < len(k); i++ {
		for j := 1; j < len(k); j++ {
			if k[j-1] > k[j] {
				k[j-1], k[j] = k[j], k[j-1]
			}
		}
	}
	return string(k)
}

func insert(s string) string {
	for i := 0; i <= len(s)-2; i++ {
		if s[i:i+2] == "-i" {
			k := equally(s) + 1
			return s[k:]
		}
	}
	return ""
}

func equally(s string) int {
	for i := 0; i <= len(s)-2; i++ {
		if s[i:i+1] == "=" {
			return i
		}
	}
	return 0
}
