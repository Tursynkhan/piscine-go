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
	if nb <= 1 {
		return 2
	}
	for i := 2; !IsPrime1(i); i++ {
		return i
	}
	return 0
}
