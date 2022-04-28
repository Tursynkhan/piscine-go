package piscine

func TrimAtoi(s string) int {
	var r []rune
	var n int
	check := "positive"
	if s == "" {
		n = 0
	} else {
		for _, i := range s {
			if i > 47 && i < 58 || i == 45 {
				r = append(r, i)
			}
		}
	}
	if len(r) != 0 {
		if r[0] == 45 {
			check = "negative"
		}
	}
	for i := 0; i < len(r); i++ {
		if r[i] > 47 && r[i] < 58 {
			n = n*10 + int(r[i]-'0')
		}
	}
	if check == "negative" {
		n = n * -1
	}
	return n
}
