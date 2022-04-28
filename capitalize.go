package piscine

func check(s rune) bool {
	for i := 0; i < 1; i++ {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
			return true
		}
	}
	return false
}

func Capitalize(s string) string {
	str := []rune(s)
	flag := true
	for i := 0; i < len(str); i++ {
		if check(str[i]) && flag {
			if str[i] >= 'a' && s[i] <= 'z' {
				str[i] = str[i] - 32
			}
			flag = false
		} else if str[i] >= 'A' && s[i] <= 'Z' {
			str[i] = str[i] + 32
		} else if !(check(str[i])) {
			flag = true
		}
	}
	return string(str)
}
