package piscine

func IsPrime1(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	prime := 1
	for i := 2; i < nb; i++ {
		if (nb % i) == 0 {
			return false
		}
	}
	if prime == 1 {
		return true
	} else {
		return false
	}
}

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime1(i) {
			return i
		}
	}
}
