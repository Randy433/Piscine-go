package piscine

func ListReverse(l *List) {
	var Prev *NodeL
	var Next *NodeL
	current := l.Head
	for current != nil {
		Next = current.Next
		current.Next = Prev
		Prev = current
		current = Next
	}
	l.Head, l.Tail = l.Tail, l.Head
}

// Write a function ListReverse that reverses the order of the elements in the list l while using the structure List.
