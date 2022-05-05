package piscine

func ActiveBits(n int) int {
	count := 0
	for n >= 1 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
