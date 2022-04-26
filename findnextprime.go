package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for i := nb + 1; i < 1000000; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
