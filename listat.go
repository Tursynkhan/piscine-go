package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var count int
	if pos == 0 {
		return l
	}
	for temp := l; temp.Next != nil; temp = temp.Next {
		count++
		if count == pos {
			return temp.Next
		}
	}
	return nil
}
