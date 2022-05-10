package piscine

func Unmatch(a []int) int {
	var count int = 0
	var res int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 == 0 {
			count = 0
			res = -1
		} else {
			res = a[i]
			return res
		}
	}
	return res
}
