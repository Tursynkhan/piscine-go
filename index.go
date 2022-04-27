package piscine

func Index(s string, toFind string) int {
	char := []rune(s)
	find := []rune(toFind)
	for i := 0; i < len(char)-len(toFind); i++ {
		if toFind == s[i:i+len(find)] {
			return i
		}
	}
	return -1
}
