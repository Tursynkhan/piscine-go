package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascend := false
	descend := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascend = true
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			descend = true
		}
	}
	return !(ascend) || !(descend)
}
