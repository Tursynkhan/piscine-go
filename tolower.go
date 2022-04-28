package piscine

func ToLower(s string) string {
	result := []rune(s)
	for index, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			result[index] = letter + 32
		}
	}
	return string(result)
}
