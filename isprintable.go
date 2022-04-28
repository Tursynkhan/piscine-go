package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter < 65 {
			return false
		} else if letter > 90 && letter < 97 {
			return false
		} else if letter > 122 {
			return false
		}
	}
	return true
}
