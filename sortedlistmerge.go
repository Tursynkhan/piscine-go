package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	iterator := n2
	for iterator != nil {
		listPushBack(n1, iterator.Data)
		iterator = iterator.Next
	}
	for value := n1; value != nil; value = value.Next {
		for temp := value.Next; temp != nil; temp = temp.Next {
			if value.Data > temp.Data {
				value.Data, temp.Data = temp.Data, value.Data
			}
		}
	}
	return n1
}
