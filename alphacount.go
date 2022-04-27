package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			counter++
		}
	}
	return counter
}
