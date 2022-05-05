package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, value := range *ptr {
		if value != "" {
			count++
		}
	}
	var newSlice []string
	for _, value := range *ptr {
		if value != "" {
			newSlice = append(newSlice, value)
		}
	}
	*ptr = newSlice
	return count
}
