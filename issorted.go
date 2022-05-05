package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := 1; j < len(a)-1; j++ {
			if a[0] < a[len(a)-1] {
				if f(a[j], a[i]) <= 0 {
					return false
				}
			}
			if a[0] > a[len(a)-1] {
				if f(a[i], a[j]) <= 0 {
					return false
				}
			}
		}
	}
	return true
}
