package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for value := l; value != nil; value = value.Next {
		for temp := value.Next; temp != nil; temp = temp.Next {
			if value.Data > temp.Data {
				value.Data, temp.Data = temp.Data, value.Data
			}
		}
	}
	return l
}
