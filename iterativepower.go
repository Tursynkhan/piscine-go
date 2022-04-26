package piscine

func IterativePower(n int, power int) int {
	result := 1
	if power < 0 || power > 40 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := 0; i < power; i++ {
			result *= n
		}
		return result
	}
}
