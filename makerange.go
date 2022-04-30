package piscine

func MakeRange(min, max int) []int {
	var temp []int
	if min > max {
		return temp
	}

	array := make([]int, min, max)
	for i := 0; i < max-min; i++ {
		array[i] = i + min
	}

	return array
}
