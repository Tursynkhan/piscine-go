package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = l.Tail.Next
	}
}
