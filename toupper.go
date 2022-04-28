package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for index, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			result[index] = letter - 32
		}
	}
	return string(result)
}
