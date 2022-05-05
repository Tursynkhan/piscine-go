package piscine

func Abort(a, b, c, d, e int) int {
	var runes []int
	runes = append(runes, a)
	runes = append(runes, b)
	runes = append(runes, c)
	runes = append(runes, d)
	runes = append(runes, e)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return runes[2]
}
