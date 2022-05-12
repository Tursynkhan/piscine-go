package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for temp := l.Head; temp != nil; temp = temp.Next {
		if comp(temp, ref) {
			return &ref
		}
	}
	return &ref
}
