package piscine

func RecursivePower(n int, power int) int {
	if power < 0 || power > 40 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return n * RecursivePower(n, power-1)
	}
}
