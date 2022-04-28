package piscine

func IsNumeric(s string) bool {
	for _, letter := range s {
		if letter < 48 {
			return false
		} else if letter > 57 {
			return false
		}
	}
	return true
}
