package piscine

func SplitWhiteSpaces(s string) []string {
	var letter []rune
	var str []string
	for i, value := range s {
		if value != 32 {
			letter = append(letter, value)
		}
		if value == 32 || i == len(s)-1 {
			if len(letter) > 0 {
				str = append(str, string(letter))
				letter = []rune{}
			}
		}
	}
	return str
}
