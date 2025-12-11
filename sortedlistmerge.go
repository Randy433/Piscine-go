package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Handle empty list cases
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	// Determine the head of the merged list
	var head *NodeI
	if n1.Data <= n2.Data {
		head = n1
		n1 = n1.Next
	} else {
		head = n2
		n2 = n2.Next
	}

	// Merge the remaining nodes
	current := head
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes from whichever list is not exhausted
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return head
}
