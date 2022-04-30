package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	i := nb
	for {
		flag := 1
		j := 2
		for j < i/2+1 {
			if i%j == 0 {
				flag = 0
				break
			}
			j++
		}
		if flag == 1 {
			break
		}
		i++
	}
	return i
}
