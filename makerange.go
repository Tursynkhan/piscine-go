package piscine

func MakeRange(min, max int) []int {
	var temp []int
	if min > max {
		return temp
	}

	array := make([]int, max-min)
	for i := 0; i < min; i++ {
		array[i] = i + min
	}

	return array
}
