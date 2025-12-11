package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	wnode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = wnode
	} else {
		wnode.Next = l.Head
		l.Head = wnode
	}
}

// Write a function ListPushFront that inserts a new element NodeL at the beginning of the list l while using the structure List
