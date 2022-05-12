package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listPushBack(l, data_ref)
	for value := l; value != nil; value = value.Next {
		for temp := value.Next; temp != nil; temp = temp.Next {
			if value.Data > temp.Data {
				value.Data, temp.Data = temp.Data, value.Data
			}
		}
	}
	return l
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
