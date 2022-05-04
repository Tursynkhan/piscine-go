package piscine

func Any(f func(string) bool, a []string) bool {
	b := true
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			b = true
			break
		} else {
			b = false
		}
	}
	return b
}
