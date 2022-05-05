package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for index, letter := range s {
		if (letter >= 'A' && letter <= 'L') || (letter >= 'a' && letter <= 'l') {
			runes[index] = letter + 14
		} else if (letter >= 'M' && letter <= 'Z') || (letter >= 'm' && letter <= 'z') {
			runes[index] = letter - 12
		}
	}
	result := ""
	for i := 0; i < len(runes); i++ {
		result = string(runes)
	}
	return string(result)
}
