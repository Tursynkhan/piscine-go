package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 39 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		if i < 0 || i > 12 {
			return 0
		} else {
			nb *= i
		}
	}
	return nb
}
