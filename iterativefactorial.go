package piscine

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			if i < 0 || i > 40 {
				return 0
			} else {
				factorial *= i
			}
		}
	}
	return factorial
}
