package piscine

func ListSize(l *List) int {
	n := l.Head
	counter := 0
	for n != nil {
		counter++
		n = n.Next
	}
	return counter
}
