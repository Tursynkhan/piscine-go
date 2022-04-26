package piscine

func Sqrt(n int) int {
	if n < 0 {
		return 0
	} else {
		for i := 0; i <= n; i++ {
			if i*i == n {
				return i
			}
		}
		return 0
	}
}
