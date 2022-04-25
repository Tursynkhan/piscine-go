package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			nb *= i
		}
		return nb
	}
}
