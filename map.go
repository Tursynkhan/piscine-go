package piscine

func Map(f func(int) bool, a []int) []bool {
	length := len(a)
	arr := make([]bool, length)
	for i := 0; i < len(a); i++ {
		arr[i] = f(a[i])
	}
	return arr
}
