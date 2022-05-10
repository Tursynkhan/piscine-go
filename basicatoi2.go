package piscine

func BasicAtoi2(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			s = ""
		}
	}
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
